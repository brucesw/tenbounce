package model

import "time"

type StatsSummary struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`

	Stats []Stat `json:"stats"`
}

type MiniPoint struct {
	Value     PointValue `json:"pointValue"`
	Timestamp time.Time  `json:"timestamp"`
}

type Stat struct {
	PointTypeID   PointTypeID   `json:"pointTypeID"`
	PointTypeName PointTypeName `json:"pointTypeName"`
	Values        []MiniPoint   `json:"values"`
}
