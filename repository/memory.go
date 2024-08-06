package repository

import (
	"encoding/json"
	"fmt"
	"tenbounce/model"
	"time"

	"github.com/google/uuid"

	_ "embed"
)

//go:embed users.json
var hardcodedUsers_bytes []byte

// TODO(bruce): XXX
var hardcodedUsers []model.User

// TODO(bruce): temp
var BSWUserID string = "550e8400-e29b-41d4-a716-446655440000"
var DTherrUserID string = "123e4567-e89b-12d3-a456-426614174000"

type Memory struct {
	points     []model.Point
	users      []model.User
	pointTypes []model.PointType
}

func NewMemoryRepository() *Memory {
	// TODO(bruce): embed initial points and other things
	var points = []model.Point{
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
		{
			ID:          uuid.NewString(),
			Timestamp:   time.Date(2024, time.July, 23, 10, 40, 0, 0, time.UTC),
			UserID:      BSWUserID,
			PointTypeID: "0d1b30ef-00d4-41d6-8581-b8d554752816",
			Value:       18,
		},
		{
			ID:          uuid.NewString(),
			Timestamp:   time.Date(2024, time.August, 1, 10, 20, 0, 0, time.UTC),
			UserID:      DTherrUserID,
			PointTypeID: "4e4b2b1c-5063-425a-a409-71b431068f78",
			Value:       20.21,
		},
		{
			ID:          uuid.NewString(),
			Timestamp:   time.Date(2024, time.August, 2, 10, 30, 0, 0, time.UTC),
			UserID:      DTherrUserID,
			PointTypeID: "4e4b2b1c-5063-425a-a409-71b431068f78",
			Value:       21,
		},
		{
			ID:          uuid.NewString(),
			Timestamp:   time.Date(2024, time.August, 3, 10, 40, 0, 0, time.UTC),
			UserID:      DTherrUserID,
			PointTypeID: "0d1b30ef-00d4-41d6-8581-b8d554752816",
			Value:       19.00,
		},
	}

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

	return &Memory{
		points:     points,
		users:      hardcodedUsers,
		pointTypes: pointTypes,
	}
}

func (r *Memory) GetUser(userID string) (model.User, error) {
	for _, user := range r.users {
		if user.ID == userID {
			return user, nil
		}
	}

	return model.User{}, fmt.Errorf("user '%s' not found", userID)
}

func (r *Memory) ListPoints(userID string) ([]model.Point, error) {
	var points = []model.Point{}

	for _, point := range r.points {
		if point.UserID == userID {
			points = append(points, point)
		}
	}

	return points, nil
}

func (r *Memory) SavePoint(p *model.Point) error {
	p.ID = uuid.NewString()
	fmt.Printf("saved to db: %+v", p)

	r.points = append(r.points, *p)

	return nil
}

func (r *Memory) ListPointTypes() ([]model.PointType, error) {
	return r.pointTypes, nil
}

func init() {
	var err = json.Unmarshal(hardcodedUsers_bytes, &hardcodedUsers)
	if err != nil {
		panic(fmt.Errorf("unmarshal hardcoded users %w", err))
	}
}
