package model

// TODO(bruce): am I going overboard with these types?
type PointTypeID string
type PointTypeName string

type PointType struct {
	ID PointTypeID `json:"id"`

	Name PointTypeName `json:"name"`
}
