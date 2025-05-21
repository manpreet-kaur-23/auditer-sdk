package constants

import "time"

type AuditLog struct {
	RecordChanges map[string][]interface{} `json:"record_changes"`
	Actions       []string                 `json:"actions"`
	ResourceID    int                      `json:"resource_id"`
	Resource      string                   `json:"resource"`
	Event         string                   `json:"event"`
	Record        map[string]interface{}   `json:"record"`
	User          User                     `json:"user"`
	Latitude      float64                  `json:"lat"`
	Longitude     float64                  `json:"lng"`
	Timestamp     time.Time                `json:"timestamp"`
	Service       string                   `json:"service"`
	MetaData      map[string]interface{}   `json:"meta_data"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
