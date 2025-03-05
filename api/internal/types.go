package internal

type Request struct {
	Method string       `json:"method"`
	Data   *RequestData `json:"data,omitempty"`
}

type RequestData struct {
	UserId    string         `json:"user_id"`               // user id
	UserIp    string         `json:"user_ip,omitempty"`     // ip address
	Severity  string         `json:"severity,omitempty"`    // info, warning, error, critical etc.
	Action    string         `json:"action"`                // action name such as login, logout etc.
	Content   map[string]any `json:"content,omitempty"`     // content
	Agent     string         `json:"user_agent,omitempty"`  // user agent
	Start     string         `json:"start_date,omitempty"`  // start date
	End       string         `json:"ending_date,omitempty"` // end date
	Page      int            `json:"page,omitempty"`        // current page
	Limit     int            `json:"limit,omitempty"`       // records per page
	Order     string         `json:"order,omitempty"`       // column name
	Direction string         `json:"direction,omitempty"`   // asc or desc
}

type Response struct {
	Code    int            `json:"code"`              // 200, 400, 401, 403, 404, 500 etc.
	Status  string         `json:"status"`            // success, error etc.
	Message string         `json:"message"`           // message
	Records int            `json:"records,omitempty"` // total records
	Page    int            `json:"page,omitempty"`    // current page
	Total   int            `json:"total,omitempty"`   // total pages
	Data    []ResponseData `json:"data,omitempty"`    // response data
}

type ResponseData struct {
	UserId    string `json:"user_id,omitempty"`    // user id
	UserIp    string `json:"user_ip,omitempty"`    // ip address
	Severity  string `json:"severity,omitempty"`   // info, warning, error, critical etc.
	Action    string `json:"action,omitempty"`     // action name such as login, logout etc.
	Content   string `json:"content,omitempty"`    // content
	Agent     string `json:"user_agent,omitempty"` // user agent
	Timestamp string `json:"timestamp,omitempty"`  // timestamp
}
