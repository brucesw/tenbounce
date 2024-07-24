package api

import (
	"fmt"
	"net/http"
	"tenbounce/model"

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

	PointTypeID model.PointTypeID `json:"pointTypeID"`

	Value model.PointValue `json:"value"`
}

func NewCreatePointResponse(p model.Point) (CreatePointResponse, error) {
	var cpr = CreatePointResponse{
		ID:          p.ID,
		PointTypeID: p.PointTypeID,
		Value:       p.Value,
	}

	return cpr, nil
}

func (cpb CreatePointBody) Point() (model.Point, error) {
	var point = model.Point{
		PointTypeID: cpb.PointTypeID,
		Value:       cpb.Value,
	}

	return point, nil
}

// TODO(bruce): use store
func SavePointToDB(p *model.Point) error {
	p.ID = uuid.NewString()
	fmt.Printf("saved to db: %+v", p)

	return nil
}

func createPoint(c echo.Context) error {
	var pointBody = &CreatePointBody{}

	if err := c.Bind(pointBody); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	} else if point, err := pointBody.Point(); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else if err = SavePointToDB(&point); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else if cpr, err := NewCreatePointResponse(point); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else {
		return c.JSON(http.StatusOK, cpr)
	}
}

func GetPointsFromDB() ([]model.Point, error) {
	var points = []model.Point{
		{
			ID:          uuid.NewString(),
			PointTypeID: "compulsory",
			Value:       20.21,
		},
		{
			ID:          uuid.NewString(),
			PointTypeID: "compulsory",
			Value:       21,
		},
		{
			ID:          uuid.NewString(),
			PointTypeID: "optional",
			Value:       19.00,
		},
	}

	return points, nil
}

type ListPointsResponse []model.Point

func listPoints(c echo.Context) error {
	if points, err := GetPointsFromDB(); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else {
		return c.JSON(http.StatusOK, ListPointsResponse(points))
	}
}
