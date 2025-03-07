package database

import (
	"fmt"
	"math"
	"os"

	"github.com/aticish/log-service/internal"
	"github.com/gofiber/fiber/v2"
)

func Read(data map[string]any, request *internal.RequestData) (*internal.Response, error) {

	// Connect to database
	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Get data
	var (
		ip       string = data["ip"].(string)
		user     string = data["user"].(string)
		severity string = data["severity"].(string)
		action   string = data["action"].(string)
		start    string = data["start"].(string)
		end      string = data["end"].(string)
		limit    string = data["limit"].(string)
		order    string = data["order"].(string)
	)

	// Prepare queries
	count := "SELECT COUNT() AS total FROM " + os.Getenv("CLICKHOUSE_TABLE") + " WHERE 1=1 " + ip + user + severity + action + start + end
	query := "SELECT * FROM " + os.Getenv("CLICKHOUSE_TABLE") + " WHERE 1=1 " + ip + user + severity + action + start + end + order + limit

	// Get total count
	row := conn.QueryRow(count)
	var total int
	if err = row.Scan(&total); err != nil {
		return nil, fmt.Errorf("failed to get total logs: %v", err)
	}

	// Pagination
	var currentLimit int
	if request.Limit < 1 {
		currentLimit = 1000
	} else {
		currentLimit = min(request.Limit, 10000)
	}

	var currentPage int
	currentPage = max(request.Page, 1)

	if total == 0 {
		return &internal.Response{
			Code:    fiber.StatusOK,
			Status:  "success",
			Message: internal.LogNotFoundMessage,
			Records: total,
			Total:   1,
			Page:    currentPage,
		}, nil
	}

	// Fetch data
	rows, err := conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch logs: %v", err)
	}
	defer rows.Close()

	// Calculate total pages
	totalPages := int(math.Ceil(float64(total) / float64(currentLimit)))
	var responseDatas []internal.ResponseData

	for rows.Next() {
		var responseData internal.ResponseData

		if err = rows.Scan(&responseData.UserId, &responseData.Severity, &responseData.UserIp, &responseData.Action, &responseData.Content, &responseData.Agent, &responseData.Timestamp); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		responseDatas = append(responseDatas, responseData)
	}

	return &internal.Response{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: internal.LogRetrievedMessage,
		Records: total,
		Total:   totalPages,
		Page:    currentPage,
		Data:    responseDatas,
	}, nil
}
