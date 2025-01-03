package models

type EmployeeRate struct {
	ID           int    `json:"id" db:"id"`
	FirstName    string `json:"first_name" db:"first_name"`
	LastName     string `json:"last_name" db:"last_name"`
	PositionID   int    `json:"position_id" db:"position_id"`
	PositionName string `json:"position_name" db:"position_name"`
	S025         int    `json:"s025" db:"s025"`
	S050         int    `json:"s050" db:"s050"`
	S075         int    `json:"s075" db:"s075"`
	S100         int    `json:"s100" db:"s100"`
	Partial      int    `json:"partial" db:"partial"`
}
