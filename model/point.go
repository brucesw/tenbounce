package model

type PointTypeID string
type PointValue float64

type Point struct {
	ID string `json:"id"`

	PointTypeID PointTypeID `json:"pointTypeID"`

	Value PointValue `json:"value"`
}
