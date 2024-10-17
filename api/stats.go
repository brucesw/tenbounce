package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tenbounce/model"

	"github.com/labstack/echo/v4"
)

func statsRoutez(g *echo.Group, h HandlerClx) {
	g.GET("/summary", h.getStatsSummary)
	g.POST("/summary_gpt", h.getStatsSummaryGPT)

}

func (h HandlerClx) getStatsSummary(c echo.Context) error {
	var ctx = c.Request().Context()

	if userID, err := contextUserID(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("context user id: %w", err))
		// TODO(bruce): confirm creator user has permission to create points for user
	} else if _, err := h.repository.GetUser(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, "get user")
	} else if statsSummary, err := h.repository.GetStatsSummary(); err != nil {
		return c.JSON(http.StatusInternalServerError, "get stats summary")
	} else {
		return c.JSON(http.StatusOK, statsSummary)
	}
}

type StatsSummaryGPT struct {
	UserID      string            `json:"userID"`
	PointTypeID model.PointTypeID `json:"pointTypeID"`
	Summary     string            `json:"summary"`
}

func generateSummary(statsSummaries []model.StatsSummary) ([]StatsSummaryGPT, error) {
	var statsSummaryGPTs []StatsSummaryGPT

	// for _, statsSummary := range statsSummaries {
	// 	var statsSummaryGPT = StatsSummaryGPT{
	// 		UserID:      statsSummary.UserID,
	// 		PointTypeID: statsSummary.PointTypeID,
	// 	}
	// }

	var summaryText = `
	[
  {
    "pointTypeID": "4e4b2b1c-5063-425a-a409-71b431068f78",
    "userID": "550e8400-e29b-41d4-a716-446655440000",
    "summary": "Bruce Szudera Wienand's Compulsory Routine shows a fluctuating trend. The values initially dropped from 21.01 to 18.15, but then rose to 24.15 before stabilizing at 21. Overall sentiment is mixed, with improvements toward the end but high variability."
  },
  {
    "pointTypeID": "0d1b30ef-00d4-41d6-8581-b8d554752816",
    "userID": "550e8400-e29b-41d4-a716-446655440000",
    "summary": "Bruce Szudera Wienand's Optional Routine shows a slight decline, dropping from 19 to 18 over two days. The overall sentiment is slightly negative due to the downward trend."
  },
  {
    "pointTypeID": "4e4b2b1c-5063-425a-a409-71b431068f78",
    "userID": "123e4567-e89b-12d3-a456-426614174000",
    "summary": "Derek Therrien's Compulsory Routine shows a consistent improvement, rising from 20.21 to 21 over two days. The overall sentiment is positive with gradual improvement."
  },
  {
    "pointTypeID": "0d1b30ef-00d4-41d6-8581-b8d554752816",
    "userID": "123e4567-e89b-12d3-a456-426614174000",
    "summary": "Derek Therrien's Optional Routine has only one recorded value of 19, so no trend can be determined, but the sentiment is neutral."
  },
  {
    "pointTypeID": null,
    "userID": "987fbc97-4bed-5078-889f-8c6e44d66b00",
    "summary": "Lourens Willekes has no stats recorded, so no trend or sentiment can be determined."
  },
  {
    "pointTypeID": null,
    "userID": "d6f0bf56-0abb-4278-9e2c-c3e0bfc18c1d",
    "summary": "Nadia has no stats recorded, so no trend or sentiment can be determined."
  }
]
`
	if err := json.Unmarshal([]byte(summaryText), &statsSummaryGPTs); err != nil {
		return nil, fmt.Errorf("unmarshal summary text: %w", err)
	}

	return statsSummaryGPTs, nil
}

func (h HandlerClx) getStatsSummaryGPT(c echo.Context) error {
	var ctx = c.Request().Context()
	var statsSummaries []model.StatsSummary

	if userID, err := contextUserID(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("context user id: %w", err))
		// TODO(bruce): confirm creator user has permission to create points for user
	} else if _, err := h.repository.GetUser(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, "get user")
		// } else if err := c.Bind(&statsSummaries); err != nil {
		// 	return c.JSON(http.StatusInternalServerError, "bind stats summary")
	} else if gptSummaries, err := generateSummary(statsSummaries); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("generate summary: %w", err))
	} else {
		fmt.Println(gptSummaries)
		return c.JSON(http.StatusOK, gptSummaries)
	}
}
