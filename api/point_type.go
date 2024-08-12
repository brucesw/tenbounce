package api

import (
	"fmt"
	"net/http"
	"tenbounce/model"

	"github.com/labstack/echo/v4"
)

func pointTypeRoutes(g *echo.Group, h HandlerClx) {
	var pointTypeRoutes = g.Group("/point_types")

	pointTypeRoutes.GET("", h.listPointTypes)
	pointTypeRoutes.POST("", h.createPointType)

}

type CreatePointTypeBody struct {
	Name model.PointTypeName `json:"name"`
}

func (cptb CreatePointTypeBody) PointType() (model.PointType, error) {
	var pointType = model.PointType{
		Name: cptb.Name,
	}

	return pointType, nil
}

// TODO(bruce): document
// TODO(bruce): responses
func (h HandlerClx) createPointType(c echo.Context) error {
	var ctx = c.Request().Context()
	var createPointTypeBody = &CreatePointTypeBody{}

	if err := c.Bind(createPointTypeBody); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid point type body")
	} else if userID, err := contextUserID(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("context user id: %w", err))
		// TODO(bruce): confirm creator user has permission to create point types
	} else if _, err := h.repository.GetUser(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, "get user")
	} else if pointType, err := createPointTypeBody.PointType(); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("create point type body point type: %w", err))
	} else if err = h.repository.CreatePointType(&pointType); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("create point type: %w", err))
	} else {
		return c.JSON(http.StatusOK, nil)
	}
}

// TODO(bruce): document
// TODO(bruce): responses
func (h HandlerClx) listPointTypes(c echo.Context) error {
	var ctx = c.Request().Context()

	if userID, err := contextUserID(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("context user id: %w", err))
	} else if _, err := h.repository.GetUser(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("get user: %w", err))
	} else if pointTypes, err := h.repository.ListPointTypes(); err != nil {
		return c.JSON(http.StatusInternalServerError, "get point types from db")
	} else {
		return c.JSON(http.StatusOK, pointTypes)
	}
}
