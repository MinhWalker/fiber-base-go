package request

type CreateContestRequest struct {
	Name    string   `json:"name" validate:"required"`
	Classes []string `json:"classes"`
	Grades  []string `json:"grades"`
}
