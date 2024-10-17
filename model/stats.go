package model

type StatsSummary struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`

	Stats []Stat `json:"stats"`
}

type Stat struct {
	PointTypeID   PointTypeID   `json:"pointTypeID"`
	PointTypeName PointTypeName `json:"pointTypeName"`
	Value         PointValue    `json:"pointValue"`
}
