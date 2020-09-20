package model

// EmployeeSSO ..
type EmployeeSSO struct {
	EmployeeID int    `json:"employee_id"`
	SSOToken   string `json:"sso_token"`
}
