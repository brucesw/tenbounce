package api

import (
	"fmt"
	"net/http"
	"tenbounce/model"
	"tenbounce/util"

	"time"

	"github.com/labstack/echo/v4"
)

func pointRoutes(g *echo.Group, h HandlerClx) {
	var pointRoutes = g.Group("/points")

	pointRoutes.POST("", h.createPoint)
	pointRoutes.GET("", h.listPoints)
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

func (cpb CreatePointBody) Point(user model.User) (model.Point, error) {
	var point = model.Point{
		// ID set downstream
		Timestamp:   time.Now(), // TODO(bruce): use nower
		UserID:      user.ID,
		PointTypeID: cpb.PointTypeID,
		Value:       cpb.Value,
	}

	return point, nil
}

// TODO(bruce): document
// TODO(bruce): responses
func (h HandlerClx) createPoint(c echo.Context) error {
	var ctx = c.Request().Context()
	var pointBody = &CreatePointBody{}

	if err := c.Bind(pointBody); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid point body")
	} else if userID, err := contextUserID(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("context user id: %w", err))
	} else if user, err := h.repository.GetUser(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, "get user")
	} else if point, err := pointBody.Point(user); err != nil {
		return c.JSON(http.StatusInternalServerError, "pointbody point")
	} else if pointTypes, err := h.repository.ListPointTypes(); err != nil {
		return c.JSON(http.StatusInternalServerError, "get point types from db")
	} else if err = validPointTypeID(point, pointTypes); err != nil {
		return c.JSON(http.StatusInternalServerError, "valid point type id")
	} else if err = h.repository.CreatePoint(&point); err != nil {
		return c.JSON(http.StatusInternalServerError, "save point to db")
	} else if cpr, err := NewCreatePointResponse(point); err != nil {
		return c.JSON(http.StatusInternalServerError, "new create point response")
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
func (h HandlerClx) listPoints(c echo.Context) error {
	var ctx = c.Request().Context()

	if userID, err := contextUserID(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("context user id: %w", err))
	} else if user, err := h.repository.GetUser(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, "get user")
	} else if points, err := h.repository.ListPoints(user.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else if pointTypes, err := h.repository.ListPointTypes(); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else if response, err := NewListPointsResponse(points, pointTypes); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		return c.JSON(http.StatusOK, response)
	}
}
