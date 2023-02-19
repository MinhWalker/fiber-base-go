package request

import "time"

type UpdateStudentRequest struct {
	Name     string    `json:"name,omitempty"`
	Class    string    `json:"class,omitempty"`
	Birthday time.Time `json:"birthday,omitempty"`
}
