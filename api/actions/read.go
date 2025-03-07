package actions

import (
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aticish/log-service/database"
	"github.com/aticish/log-service/internal"
)

func Read(request *internal.RequestData) (*internal.Response, error) {

	// empty request
	if request == nil {
		return nil, internal.ErrorInvalidRequest
	}

	// ip query
	ip := buildIpQuery(request.UserIp)

	// user id query
	user := buildUserQuery(request.UserId)

	// Severity query
	severity := buildSeverityQuery(request.Severity)

	// Action query
	action := buildActionQuery(request.Action)

	// Start date query
	start := buildStartDateQuery(request.Start)

	// End date query
	end := buildEndDateQuery(request.End)

	// limit query
	limit := buildLimitQuery(request.Limit, request.Page)

	// order
	order := buildOrderQuery(request.Order, request.Direction)

	data := map[string]any{
		"ip":       ip,
		"user":     user,
		"severity": severity,
		"action":   action,
		"start":    start,
		"end":      end,
		"limit":    limit,
		"order":    order,
	}

	// Read from db
	result, err := database.Read(data, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// buildOrderQuery builds the order query
func buildOrderQuery(order string, direction string) string {

	order = strings.ToLower(strings.TrimSpace(order))
	direction = strings.ToLower(strings.TrimSpace(direction))

	d := " DESC "
	if direction == "asc" || direction == "ascending" {
		d = " ASC "
	}

	if order == "timestamp" || order == "user_id" || order == "severity" || order == "user_ip" || order == "action" || order == "admin_id" {
		return " ORDER BY " + order + d
	}

	return " ORDER BY timestamp " + d

}

// buildLimitQuery builds the limit query
func buildLimitQuery(limit int, page int) string {

	var currentLimit int

	if limit < 1 {
		currentLimit = 1000
	} else if limit > 10000 {
		currentLimit = 10000
	} else {
		currentLimit = limit
	}

	var currentPage int
	if page < 1 {
		currentPage = 1
	} else {
		currentPage = page
	}

	var currentOffset = (currentPage - 1) * currentLimit

	// Prepare pagination query
	stringLimit := strconv.Itoa(currentLimit)
	stringOffset := strconv.Itoa(currentOffset)
	return " LIMIT " + stringLimit + " OFFSET " + stringOffset

}

// buildEndDateQuery builds the end date query
func buildEndDateQuery(end string) string {

	end = strings.TrimSpace(end)
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02 15:04:05")
	_, err := time.Parse("2006-01-02 15:04:05", end)

	if err != nil || end == "" {
		return " AND ( timestamp < '" + tomorrow + "' ) "
	}

	return " AND ( timestamp < '" + end + "' ) "

}

// buildStartDateQuery builds the start date query
func buildStartDateQuery(start string) string {

	start = strings.TrimSpace(start)
	oneMonthAgo := time.Now().AddDate(0, -1, 0).Format("2006-01-02 15:04:05")

	_, err := time.Parse("2006-01-02 15:04:05", start)

	if err != nil || start == "" {
		return " AND ( timestamp >= '" + oneMonthAgo + "' ) "
	}

	return " AND ( timestamp >= '" + start + "' ) "

}

// buildActionQuery builds the action query
func buildActionQuery(action string) string {

	action = strings.TrimSpace(action)

	if action == "*" || action == "any" || action == "" {
		return ""
	}

	// more than one action given
	if strings.Contains(action, ",") {
		parts := strings.Split(action, ",")
		var actions []string
		for _, single := range parts {
			trimmedPart := strings.TrimSpace(single)
			if trimmedPart != "" {
				actions = append(actions, trimmedPart)
			}
		}

		if len(actions) > 0 {
			return " AND ( action = '" + strings.Join(actions, "' OR action = '") + "' ) "
		}

		return ""
	}

	return " AND ( action = '" + action + "' ) "

}

// buildSeverityQuery builds the severity query
func buildSeverityQuery(severity string) string {

	severity = strings.TrimSpace(severity)

	if severity == "*" || severity == "any" || severity == "" {
		return ""
	}

	// more than one severity given
	if strings.Contains(severity, ",") {
		parts := strings.Split(severity, ",")
		var severities []string
		for _, single := range parts {
			trimmedPart := strings.TrimSpace(single)
			checked := internal.CheckSeverity(trimmedPart)
			if trimmedPart != "" && checked {
				severities = append(severities, trimmedPart)
			}
		}

		if len(severities) > 0 {
			return " AND ( severity = '" + strings.Join(severities, "' OR severity = '") + "' ) "
		}

		return ""
	}

	// single severity
	isvalid := internal.CheckSeverity(severity)

	if !isvalid {
		return ""
	}

	return " AND ( severity = '" + severity + "' ) "
}

// buildUserQuery builds the user query
func buildUserQuery(user string) string {

	acceptedCharacters := "[^0-9]"

	user = strings.TrimSpace(user)
	user = replaceAsteriks(user)

	// no user id given, select all users
	if user == "*" || user == "any" || user == "" {
		return ""
	}

	// more than one user id given
	if strings.Contains(user, ",") {

		parts := strings.Split(user, ",")
		var users []string
		for _, single := range parts {
			trimmedPart := strings.TrimSpace(single)
			containsInvalidCharacters, _ := regexp.MatchString(acceptedCharacters, trimmedPart)
			if trimmedPart != "" && !containsInvalidCharacters {
				users = append(users, trimmedPart)
			}
		}
		if len(users) > 0 {
			return " AND ( user_id IN ('" + strings.Join(users, "','") + "') ) "
		}

		return ""

	}

	// single user id
	containsInvalidCharacters, _ := regexp.MatchString(acceptedCharacters, user)
	if containsInvalidCharacters {
		return ""
	}

	return " AND ( user_id IN ('" + user + "') ) "

}

// buildIpQuery builds the ip query
func buildIpQuery(ip string) string {

	var query string

	ip = strings.TrimSpace(ip)

	// Replace invalid asteriks added adjointly more than two
	ip = replaceAsteriks(ip)

	// Check if the ip has invalid characters
	acceptedCharacters := "[^0-9a-zA-Z.*,]"
	containsInvalidCharacters, _ := regexp.MatchString(acceptedCharacters, ip)

	// any ip
	if ip == "*" || ip == "any" || containsInvalidCharacters {
		return ""
	}

	// Single ip
	if net.ParseIP(ip) != nil {
		return " AND ( user_ip = '" + ip + "' ) "
	}

	// Multiple ips
	if strings.Contains(ip, ",") {

		// Split the ip addresses
		ips := strings.Split(ip, ",")
		query = " AND ( "
		for i, ipOne := range ips {
			if net.ParseIP(ipOne) != nil {
				query += " user_ip = '" + ipOne + "' "
				if i != len(ips)-1 {
					query += " OR "
				}
			}
		}
		query += " ) "

		return query

	}

	// Contains asteriks
	if strings.Contains(ip, "*") {
		return " AND ( user_ip LIKE '" + strings.ReplaceAll(ip, "*", "%") + "' ) "
	}

	return " AND ( user_ip LIKE '%" + ip + "%' ) "
}

// replaceAsteriks replaces invalid asteriks added adjointly more than two
func replaceAsteriks(value string) string {
	invalidAsteriks := regexp.MustCompile(`\*{2,}`)
	value = invalidAsteriks.ReplaceAllString(value, "*")
	return value
}
