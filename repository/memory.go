package repository

import (
	"encoding/json"
	"fmt"
	"tenbounce/model"
	"tenbounce/util"
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

	nower util.Nower
}

func NewMemoryRepository(nower util.Nower) *Memory {
	// TODO(bruce): embed initial points and other things
	var points = []model.Point{
		{
			ID:              uuid.NewString(),
			Timestamp:       time.Date(2024, time.July, 20, 10, 20, 0, 0, time.UTC),
			UserID:          BSWUserID,
			PointTypeID:     "4e4b2b1c-5063-425a-a409-71b431068f78",
			Value:           20.21,
			CreatedByUserID: BSWUserID,
		},
		{
			ID:              uuid.NewString(),
			Timestamp:       time.Date(2024, time.July, 20, 10, 20, 0, 0, time.UTC),
			UserID:          BSWUserID,
			PointTypeID:     "4e4b2b1c-5063-425a-a409-71b431068f78",
			Value:           21.01,
			CreatedByUserID: BSWUserID,
		},
		{
			ID:              uuid.NewString(),
			Timestamp:       time.Date(2024, time.July, 20, 11, 20, 0, 0, time.UTC),
			UserID:          BSWUserID,
			PointTypeID:     "4e4b2b1c-5063-425a-a409-71b431068f78",
			Value:           19.10,
			CreatedByUserID: BSWUserID,
		},
		{
			ID:              uuid.NewString(),
			Timestamp:       time.Date(2024, time.July, 20, 12, 20, 0, 0, time.UTC),
			UserID:          BSWUserID,
			PointTypeID:     "4e4b2b1c-5063-425a-a409-71b431068f78",
			Value:           18.15,
			CreatedByUserID: BSWUserID,
		},
		{
			ID:              uuid.NewString(),
			Timestamp:       time.Date(2024, time.July, 20, 13, 20, 0, 0, time.UTC),
			UserID:          BSWUserID,
			PointTypeID:     "4e4b2b1c-5063-425a-a409-71b431068f78",
			Value:           22.21,
			CreatedByUserID: BSWUserID,
		},
		{
			ID:              uuid.NewString(),
			Timestamp:       time.Date(2024, time.July, 20, 14, 20, 0, 0, time.UTC),
			UserID:          BSWUserID,
			PointTypeID:     "4e4b2b1c-5063-425a-a409-71b431068f78",
			Value:           24.15,
			CreatedByUserID: BSWUserID,
		},
		{
			ID:              uuid.NewString(),
			Timestamp:       time.Date(2024, time.July, 21, 15, 30, 0, 0, time.UTC),
			UserID:          BSWUserID,
			PointTypeID:     "4e4b2b1c-5063-425a-a409-71b431068f78",
			Value:           21,
			CreatedByUserID: BSWUserID,
		},
		{
			ID:              uuid.NewString(),
			Timestamp:       time.Date(2024, time.July, 22, 10, 40, 0, 0, time.UTC),
			UserID:          BSWUserID,
			PointTypeID:     "0d1b30ef-00d4-41d6-8581-b8d554752816",
			Value:           19.00,
			CreatedByUserID: BSWUserID,
		},
		{
			ID:              uuid.NewString(),
			Timestamp:       time.Date(2024, time.July, 23, 10, 40, 0, 0, time.UTC),
			UserID:          BSWUserID,
			PointTypeID:     "0d1b30ef-00d4-41d6-8581-b8d554752816",
			Value:           18,
			CreatedByUserID: BSWUserID,
		},
		{
			ID:              uuid.NewString(),
			Timestamp:       time.Date(2024, time.August, 1, 10, 20, 0, 0, time.UTC),
			UserID:          DTherrUserID,
			PointTypeID:     "4e4b2b1c-5063-425a-a409-71b431068f78",
			Value:           20.21,
			CreatedByUserID: BSWUserID,
		},
		{
			ID:              uuid.NewString(),
			Timestamp:       time.Date(2024, time.August, 2, 10, 30, 0, 0, time.UTC),
			UserID:          DTherrUserID,
			PointTypeID:     "4e4b2b1c-5063-425a-a409-71b431068f78",
			Value:           21,
			CreatedByUserID: BSWUserID,
		},
		{
			ID:              uuid.NewString(),
			Timestamp:       time.Date(2024, time.August, 3, 10, 40, 0, 0, time.UTC),
			UserID:          DTherrUserID,
			PointTypeID:     "0d1b30ef-00d4-41d6-8581-b8d554752816",
			Value:           19.00,
			CreatedByUserID: BSWUserID,
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
		nower:      nower,
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

func (r *Memory) ListUsers() ([]model.User, error) {
	return r.users, nil
}

func (r *Memory) GetPoint(pointID string) (model.Point, error) {
	for _, point := range r.points {
		if point.ID == pointID {
			return point, nil
		}
	}

	return model.Point{}, ErrPointDoesNotExist
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

func (r *Memory) CreatePoint(p *model.Point) error {
	p.ID = uuid.NewString()

	r.points = append(r.points, *p)

	return nil
}

func (r *Memory) DeletePoint(pointID string) error {
	for i, point := range r.points {
		if point.ID == pointID {
			// Swap final element with found element, then drop final element off slice
			r.points[i] = r.points[len(r.points)-1]
			r.points = r.points[:len(r.points)-1]

			return nil
		}
	}

	return ErrPointDoesNotExist
}

func (r *Memory) ListPointTypes() ([]model.PointType, error) {
	return r.pointTypes, nil
}

func (r *Memory) CreatePointType(p *model.PointType) error {
	p.ID = model.PointTypeID(uuid.NewString())

	r.pointTypes = append(r.pointTypes, *p)

	return nil
}

func (r *Memory) GetStatsSummary() ([]model.StatsSummary, error) {
	var statsSummaries = []model.StatsSummary{}

	for _, user := range r.users {
		var statsSummary = model.StatsSummary{
			UserID:   user.ID,
			UserName: user.Name,
		}

		statsSummaries = append(statsSummaries, statsSummary)
	}

	for _, point := range r.points {
		for _, pointType := range r.pointTypes {
			if point.PointTypeID == pointType.ID {
				// Create a new MiniPoint for the current point
				miniPoint := model.MiniPoint{
					Value:     point.Value,
					Timestamp: point.Timestamp,
				}

				for i, statsSummary := range statsSummaries {
					if statsSummary.UserID == point.UserID {
						// Find the corresponding stat for the PointTypeID
						var found bool
						for j, stat := range statsSummaries[i].Stats {
							if stat.PointTypeID == point.PointTypeID {
								// Append the MiniPoint to the existing stat's Values
								statsSummaries[i].Stats[j].Values = append(stat.Values, miniPoint)
								found = true
								break
							}
						}

						// If stat for the PointTypeID wasn't found, create a new one
						if !found {
							newStat := model.Stat{
								PointTypeID:   point.PointTypeID,
								PointTypeName: pointType.Name,
								Values:        []model.MiniPoint{miniPoint},
							}
							statsSummaries[i].Stats = append(statsSummaries[i].Stats, newStat)
						}
					}
				}
			}
		}
	}

	return statsSummaries, nil
}

func init() {
	var err = json.Unmarshal(hardcodedUsers_bytes, &hardcodedUsers)
	if err != nil {
		panic(fmt.Errorf("unmarshal hardcoded users %w", err))
	}
}
