package api

import (
	"fmt"
	"tenbounce/model"
	"time"

	"github.com/google/uuid"
)

var pointsInDB = []model.Point{
	{
		ID:          uuid.NewString(),
		Timestamp:   time.Date(2024, time.July, 20, 10, 20, 0, 0, time.UTC),
		UserID:      BSWUserID,
		PointTypeID: "4e4b2b1c-5063-425a-a409-71b431068f78",
		Value:       20.21,
	},
	{
		ID:          uuid.NewString(),
		Timestamp:   time.Date(2024, time.July, 21, 10, 30, 0, 0, time.UTC),
		UserID:      BSWUserID,
		PointTypeID: "4e4b2b1c-5063-425a-a409-71b431068f78",
		Value:       21,
	},
	{
		ID:          uuid.NewString(),
		Timestamp:   time.Date(2024, time.July, 22, 10, 40, 0, 0, time.UTC),
		UserID:      BSWUserID,
		PointTypeID: "0d1b30ef-00d4-41d6-8581-b8d554752816",
		Value:       19.00,
	},
}

// Functions with hardcoded that will eventually make up the store
// TODO(bruce): implement store and relocate

// User

func GetUser(userID string) (model.User, error) {
	for _, hardcodedUser := range hardcodedUsers {
		if hardcodedUser.ID == userID {
			return hardcodedUser.User, nil
		}
	}

	return model.User{}, fmt.Errorf("user '%s' not found", userID)
}

// Point

func SavePointToDB(p *model.Point) error {
	p.ID = uuid.NewString()
	fmt.Printf("saved to db: %+v", p)

	// Add point to in-memory db
	pointsInDB = append(pointsInDB, *p)

	return nil
}

func GetPointsFromDB() ([]model.Point, error) {
	return pointsInDB, nil
}

// PointType

func GetPointTypesFromDB() ([]model.PointType, error) {
	var pointTypes = []model.PointType{
		{
			ID:   "4e4b2b1c-5063-425a-a409-71b431068f78",
			Name: "Compulsory Routine",
		},
		{
			ID:   "0d1b30ef-00d4-41d6-8581-b8d554752816",
			Name: "Optional Routine",
		},
		{
			ID:   "dade4383-d869-4562-a680-88cb38f9972a",
			Name: "Tenbounce",
		},
		{
			ID:   "8640f8e9-0cf6-4be4-b182-d40c21a44067",
			Name: "Ten Doubles",
		},
	}

	return pointTypes, nil
}
