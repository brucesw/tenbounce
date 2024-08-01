package model

import "time"

type PointValue float64

type Point struct {
	ID string `json:"id"`

	Timestamp time.Time `json:"timestamp"`

	UserID string `json:"userID"`

	PointTypeID PointTypeID `json:"pointTypeID"`

	Value PointValue `json:"value"`

	// TODO(bruce): need CreatedBy?
}
