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

// TODO(bruce): document
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

// TODO(bruce): document
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
