package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"tenbounce/model"

	"github.com/labstack/echo/v4"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
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

func anonymize(statsSummaries []model.StatsSummary) ([]model.StatsSummary, error) {
	var anonymousSummaries []model.StatsSummary

	for _, statsSummary := range statsSummaries {
		var anonymousSummary = model.StatsSummary{
			UserID: statsSummary.UserID,
		}

		for _, stat := range statsSummary.Stats {
			var anonymousStat = model.Stat{
				PointTypeID: stat.PointTypeID,
				Values:      stat.Values,
			}
			anonymousSummary.Stats = append(anonymousSummary.Stats, anonymousStat)
		}

		anonymousSummaries = append(anonymousSummaries, anonymousSummary)
	}

	return anonymousSummaries, nil
}

func generateSummary(ctx context.Context, statsSummaries []model.StatsSummary) ([]StatsSummaryGPT, error) {
	var statsSummaryGPTs []StatsSummaryGPT

	anonymousSummaries, err := anonymize(statsSummaries)
	if err != nil {
		return nil, fmt.Errorf("anonymize: %w", err)
	}

	var openaiClient = openai.NewClient(
		option.WithAPIKey(OpenAIAPIKey),
	)

	var prompt = "you are a gymnastics coach taking a look at an athlete's trampoline data.  higher values are better.  take the following input data and return a response that is a JSON array where each object has a pointTypeID key, userID key and summary key.  there should be one entry per pointTypeID-userID combo in the input data.  the summary should summarize any trends in the input data for that pointTypeID-userID combo including overall trend. summary should be a string and should be english.  do not cite individial data points and definitely do not cite any timestamps.  keep the summary broad and high level.   result should ONLY be a json object, no other characters indicating that it is json or otherwise."

	jsonSummaries, err := json.Marshal(anonymousSummaries)
	if err != nil {
		return nil, fmt.Errorf("json marshal summaries: %w", err)
	}

	chatCompletion, err := openaiClient.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(fmt.Sprintf("%s\n\n%s", prompt, jsonSummaries))}),
		Model: openai.F(openai.ChatModelGPT4Turbo),
	})
	if err != nil {
		return nil, fmt.Errorf("new chat completion: %w", err)
	}

	var summaryText = chatCompletion.Choices[0].Message.Content

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
	} else if err := c.Bind(&statsSummaries); err != nil {
		return c.JSON(http.StatusInternalServerError, "bind stats summary")
	} else if gptSummaries, err := generateSummary(ctx, statsSummaries); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("generate summary: %w", err))
	} else {
		fmt.Println(gptSummaries)
		return c.JSON(http.StatusOK, gptSummaries)
	}
}
