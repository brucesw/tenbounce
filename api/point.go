package api

import (
	"fmt"
	"net/http"
	"tenbounce/model"
	"tenbounce/util"

	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func pointRoutes(g *echo.Group) {
	var pointRoutes = g.Group("/points")

	pointRoutes.POST("", createPoint)
	pointRoutes.GET("", listPoints)
}

// TODO(bruce): Share types with db?
type CreatePointBody struct {
	PointTypeID model.PointTypeID `json:"pointTypeID"`

	Value model.PointValue `json:"value"`
}

type CreatePointResponse struct {
	ID string `json:"id"`

	Timestamp time.Time `json:"timestamp"`

	UserID string `json:"userID"`

	PointTypeID model.PointTypeID `json:"pointTypeID"`

	Value model.PointValue `json:"value"`
}

func NewCreatePointResponse(p model.Point) (CreatePointResponse, error) {
	var cpr = CreatePointResponse{
		ID:          p.ID,
		Timestamp:   p.Timestamp,
		UserID:      p.UserID,
		PointTypeID: p.PointTypeID,
		Value:       p.Value,
	}

	return cpr, nil
}

func (cpb CreatePointBody) Point() (model.Point, error) {
	var point = model.Point{
		Timestamp:   time.Now(), // TODO(bruce): introduce and use nower
		UserID:      BSWUserID,
		PointTypeID: cpb.PointTypeID,
		Value:       cpb.Value,
	}

	return point, nil
}

// TODO(bruce): document
// TODO(bruce): responses
func createPoint(c echo.Context) error {
	var pointBody = &CreatePointBody{}

	if err := c.Bind(pointBody); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	} else if point, err := pointBody.Point(); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else if pointTypes, err := GetPointTypesFromDB(); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else if err = validPointTypeID(point, pointTypes); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else if err = SavePointToDB(&point); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else if cpr, err := NewCreatePointResponse(point); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else {
		return c.JSON(http.StatusOK, cpr)
	}
}

// TODO(bruce): document
// TODO(bruce): test
func validPointTypeID(point model.Point, pointTypes []model.PointType) error {
	var validPointTypeIDs = []model.PointTypeID{}

	for _, pointType := range pointTypes {
		validPointTypeIDs = append(validPointTypeIDs, pointType.ID)
	}

	if !util.Contains(validPointTypeIDs, point.PointTypeID) {
		return fmt.Errorf("invalid pointTypeID: %s", point.PointTypeID)
	}

	return nil
}

type pointWithDetails struct {
	model.Point

	PointTypeName model.PointTypeName `json:"pointTypeName"`
}
type ListPointsResponse []pointWithDetails

// TODO(bruce): document
// TODO(bruce): test
// TODO(bruce): replace with join in db??
func NewListPointsResponse(points []model.Point, pointTypes []model.PointType) (ListPointsResponse, error) {
	var pointTypeIDToName = map[model.PointTypeID]model.PointTypeName{}
	for _, pointType := range pointTypes {
		pointTypeIDToName[pointType.ID] = pointType.Name
	}

	var response = []pointWithDetails{}

	for _, point := range points {
		if pointTypeName, ok := pointTypeIDToName[model.PointTypeID(point.PointTypeID)]; !ok {
			return nil, fmt.Errorf("invalid point type id %s on point %s", point.PointTypeID, point.ID)
		} else {
			var pointWithDeets = pointWithDetails{
				Point:         point,
				PointTypeName: pointTypeName,
			}

			response = append(response, pointWithDeets)
		}
	}

	return response, nil
}

// TODO(bruce): document
// TODO(bruce): responses
func listPoints(c echo.Context) error {
	if points, err := GetPointsFromDB(); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else if pointTypes, err := GetPointTypesFromDB(); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else if response, err := NewListPointsResponse(points, pointTypes); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		return c.JSON(http.StatusOK, response)
	}
}

// FUTURE STORE METHODS

// TODO(bruce): use store
func SavePointToDB(p *model.Point) error {
	p.ID = uuid.NewString()
	fmt.Printf("saved to db: %+v", p)

	return nil
}

// TODO(bruce): use store
func GetPointsFromDB() ([]model.Point, error) {
	var points = []model.Point{
		{
			ID:          uuid.NewString(),
			Timestamp:   time.Date(2024, time.April, 20, 10, 20, 0, 0, time.UTC),
			UserID:      BSWUserID,
			PointTypeID: "4e4b2b1c-5063-425a-a409-71b431068f78",
			Value:       20.21,
		},
		{
			ID:          uuid.NewString(),
			Timestamp:   time.Date(2024, time.April, 20, 10, 30, 0, 0, time.UTC),
			UserID:      BSWUserID,
			PointTypeID: "4e4b2b1c-5063-425a-a409-71b431068f78",
			Value:       21,
		},
		{
			ID:          uuid.NewString(),
			Timestamp:   time.Date(2024, time.April, 20, 10, 40, 0, 0, time.UTC),
			UserID:      BSWUserID,
			PointTypeID: "0d1b30ef-00d4-41d6-8581-b8d554752816",
			Value:       19.00,
		},
	}

	return points, nil
}
