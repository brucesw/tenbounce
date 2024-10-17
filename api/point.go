package api

import (
	"errors"
	"fmt"
	"net/http"
	"tenbounce/model"
	"tenbounce/repository"
	"tenbounce/util"

	"time"

	"github.com/labstack/echo/v4"
)

func pointRoutes(g *echo.Group, h HandlerClx) {
	var pointRoutes = g.Group("/points")

	pointRoutes.POST("", h.createPoint)
	pointRoutes.GET("", h.listPoints)
	pointRoutes.DELETE("/:id", h.deletePoint)
}

// TODO(bruce): Share types with db?
type CreatePointBody struct {
	PointTypeID model.PointTypeID `json:"pointTypeID"`

	Value model.PointValue `json:"value"`

	UserID string `json:"userID"`
}

type CreatePointResponse struct {
	ID string `json:"id"`

	Timestamp time.Time `json:"timestamp"`

	UserID string `json:"userID"`

	PointTypeID model.PointTypeID `json:"pointTypeID"`

	Value model.PointValue `json:"value"`

	CreatedByUserID string `json:"createdByUserID"`
}

func NewCreatePointResponse(p model.Point) (CreatePointResponse, error) {
	var cpr = CreatePointResponse{
		ID:              p.ID,
		Timestamp:       p.Timestamp,
		UserID:          p.UserID,
		PointTypeID:     p.PointTypeID,
		Value:           p.Value,
		CreatedByUserID: p.CreatedByUserID,
	}

	return cpr, nil
}

// TODO(bruce): rethink existence of this function
func (cpb CreatePointBody) Point(creatorUserID string, ts time.Time) (model.Point, error) {
	var point = model.Point{
		// ID set downstream
		Timestamp:       ts,
		UserID:          cpb.UserID,
		PointTypeID:     cpb.PointTypeID,
		Value:           cpb.Value,
		CreatedByUserID: creatorUserID,
	}

	return point, nil
}

// TODO(bruce): document
// TODO(bruce): responses
func (h HandlerClx) createPoint(c echo.Context) error {
	var ctx = c.Request().Context()
	var createPointBody = &CreatePointBody{}

	if err := c.Bind(createPointBody); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid point body")
	} else if creatorUserID, err := contextUserID(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("context user id: %w", err))
		// TODO(bruce): confirm creator user has permission to create points for user
	} else if creatorUser, err := h.repository.GetUser(creatorUserID); err != nil {
		return c.JSON(http.StatusInternalServerError, "get creator user")
	} else if _, err := h.repository.GetUser(createPointBody.UserID); err != nil {
		return c.JSON(http.StatusInternalServerError, "get user")
	} else if point, err := createPointBody.Point(creatorUser.ID, h.nower.Now()); err != nil {
		return c.JSON(http.StatusInternalServerError, "createPointBody point")
	} else if pointTypes, err := h.repository.ListPointTypes(); err != nil {
		return c.JSON(http.StatusInternalServerError, "get point types from db")
	} else if err = validPointTypeID(point, pointTypes); err != nil {
		return c.JSON(http.StatusInternalServerError, "valid point type id")
	} else if err = h.repository.CreatePoint(&point); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("save point to db: %w", err))
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

	CreatedByUserName string `json:"createdByUserName"`
}

type ListPointsResponse []pointWithDetails

// TODO(bruce): document
// TODO(bruce): test
// TODO(bruce): replace with join in db??
func NewListPointsResponse(points []model.Point, pointTypes []model.PointType, users []model.User) (ListPointsResponse, error) {
	var pointTypeIDToName = map[model.PointTypeID]model.PointTypeName{}
	for _, pointType := range pointTypes {
		pointTypeIDToName[pointType.ID] = pointType.Name
	}

	var userIDToName = map[string]string{}
	for _, user := range users {
		userIDToName[user.ID] = user.Name
	}

	var response = []pointWithDetails{}

	for _, point := range points {
		if pointTypeName, ok := pointTypeIDToName[model.PointTypeID(point.PointTypeID)]; !ok {
			return nil, fmt.Errorf("invalid point type id %s on point %s", point.PointTypeID, point.ID)
		} else if userName, ok := userIDToName[point.CreatedByUserID]; !ok {
			return nil, fmt.Errorf("user id %s on point %s", point.CreatedByUserID, point.ID)
		} else {
			var pointWithDeets = pointWithDetails{
				Point:             point,
				PointTypeName:     pointTypeName,
				CreatedByUserName: userName,
			}

			response = append(response, pointWithDeets)
		}
	}

	return response, nil
}

// TODO(bruce): document
func (h HandlerClx) listPoints(c echo.Context) error {
	var ctx = c.Request().Context()

	if userID, err := contextUserID(ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("context user id: %w", err))
	} else if user, err := h.repository.GetUser(userID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("get user: %w", err))
	} else if points, err := h.repository.ListPoints(user.ID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("list points: %w", err))
	} else if pointTypes, err := h.repository.ListPointTypes(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("list point types: %w", err))
	} else if users, err := h.repository.ListUsers(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("list users: %w", err))
	} else if response, err := NewListPointsResponse(points, pointTypes, users); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("new list points response: %w", err))
	} else {
		return c.JSON(http.StatusOK, response)
	}
}

func (h HandlerClx) deletePoint(c echo.Context) error {
	var ctx = c.Request().Context()

	if userID, err := contextUserID(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("context user id: %w", err))
	} else if _, err := h.repository.GetUser(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, "get user")
	} else if pointID := c.Param("id"); pointID == "" {
		return c.JSON(http.StatusBadRequest, errors.New("blank point ID"))
	} else if point, err := h.repository.GetPoint(pointID); err != nil {
		if err == repository.ErrPointDoesNotExist {
			return c.JSON(http.StatusInternalServerError, "point does not exist")
		} else {
			return c.JSON(http.StatusInternalServerError, fmt.Errorf("get point: %w", err))
		}
	} else if !(point.CreatedByUserID == userID || point.UserID == userID) {
		return c.JSON(http.StatusBadRequest, errors.New("point must belong to or be created by user"))
	} else if err = h.repository.DeletePoint(pointID); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("delete point: %w", err))
	} else {
		return c.JSON(http.StatusOK, nil)
	}
}
