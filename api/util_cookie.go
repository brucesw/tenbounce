package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

const CookieName_UserID string = "TENBOUNCE_USER_ID"

// userID_ToCookieValue hashes the input user ID and signing secret to form a cookie to be placed on the browser
// produces the input to userID_FromCookieValue
// ex: 550e8400-e29b-41d4-a716-446655440000, i3MWlSX6BQqwqFKuz+tnOFGfZFYQL8ws4jtziotcHHZMDqYLtnyicenaPL7ipJ3oQJrZpocqO41akdyz77jqdg==
// -> 550e8400-e29b-41d4-a716-446655440000|FJzgY4SioS7WZznGBxkNujMP3m4NJ_21m8erJp7JbqY=
func userID_ToCookieValue(userID, secret string) (string, error) {
	var mac = hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(userID))
	var signature = mac.Sum(nil)

	var base64EncodedSignature = base64.URLEncoding.EncodeToString(signature)

	var cookieValue = fmt.Sprintf("%s|%s", userID, base64EncodedSignature)
	if len(cookieValue) > 4096 {
		return "", fmt.Errorf("cookie value too long")
	}

	return cookieValue, nil
}

// userID_FromCookieValue validates the input cookie value against the signing secret
// and returns the user ID if it matches, errors otherwise
// ex 550e8400-e29b-41d4-a716-446655440000|FJzgY4SioS7WZznGBxkNujMP3m4NJ_21m8erJp7JbqY=, i3MWlSX6BQqwqFKuz+tnOFGfZFYQL8ws4jtziotcHHZMDqYLtnyicenaPL7ipJ3oQJrZpocqO41akdyz77jqdg==
// -> 550e8400-e29b-41d4-a716-446655440000
func userID_FromCookieValue(cookieValue, secret string) (string, error) {
	var strings = strings.Split(cookieValue, "|")
	if len(strings) != 2 {
		return "", errors.New("invalid cookie string")
	}

	var userID, haveBase64EncodedSignature string = strings[0], strings[1]

	haveSignature, err := base64.URLEncoding.DecodeString(haveBase64EncodedSignature)
	if err != nil {
		return "", fmt.Errorf("base64 decode string: %w", err)
	}

	var mac = hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(userID))
	var wantSignature = mac.Sum(nil)

	if !hmac.Equal([]byte(haveSignature), wantSignature) {
		return "", errors.New("invalid signature")
	}

	return userID, nil
}
