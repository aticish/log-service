package actions

import (
	"strconv"
	"strings"
	"time"

	"github.com/aticish/log-service/database"
	"github.com/aticish/log-service/internal"
	"github.com/gofiber/fiber/v2"
)

func Write(data *internal.RequestData) (*internal.Response, error) {

	// empty request
	if data == nil {
		return nil, internal.ErrorInvalidRequest
	}

	// Given ip address is invalid
	if data.UserIp != "unknown" && !internal.CheckIP(data.UserIp) {
		return nil, internal.ErrorInvalidUserIp
	}

	// Is user id valid
	userId, err := strconv.Atoi(data.UserId)
	if err != nil {
		return nil, internal.ErrorInvalidUserId
	}

	// Severity level
	var severity = "info"
	if internal.CheckSeverity(data.Severity) {
		severity = strings.ToLower(data.Severity)
	}

	// Action
	action := strings.ToLower(strings.TrimSpace(data.Action))
	if len(action) == 0 {
		return nil, internal.ErrorInvalidAction
	}

	// Check if the content is valid JSON
	if internal.CheckContent(data.Content) == false {
		return nil, internal.ErrorInvalidContent
	}

	// Action
	agent := strings.ToLower(strings.TrimSpace(data.Agent))
	if len(agent) == 0 {
		return nil, internal.ErrorInvalidUserAgent
	}

	currentTime := time.Now().Unix()
	err = database.Write(map[string]any{
		"user_id":   userId,
		"user_ip":   data.UserIp,
		"severity":  severity,
		"action":    action,
		"content":   data.Content,
		"agent":     agent,
		"timestamp": currentTime,
	})

	if err != nil {
		return nil, err
	}

	return &internal.Response{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: internal.LogWrittenMessage,
		Data:    nil,
	}, nil
}
