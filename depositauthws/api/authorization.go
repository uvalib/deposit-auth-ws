package api

// Authorization -- the authorization attributes from SIS
type Authorization struct {
	InboundID   string `json:"inbound_id,omitempty"`
	ID          string `json:"id,omitempty"`
	EmployeeID  string `json:"employee_id,omitempty"`
	ComputingID string `json:"computing_id,omitempty"`

	FirstName  string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`

	Career     string `json:"career,omitempty"`
	Program    string `json:"program,omitempty"`
	Plan       string `json:"plan,omitempty"`
	Degree     string `json:"degree,omitempty"`
	Department string `json:"department,omitempty"`
	Title      string `json:"title,omitempty"`

	DocType string `json:"doctype,omitempty"`
	LibraID string `json:"libra_id,omitempty"`
	Status  string `json:"status,omitempty"`

	ApprovedAt string `json:"approved_at,omitempty"`
	AcceptedAt string `json:"accepted_at,omitempty"`
	ExportedAt string `json:"exported_at,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
}

//
// end of file
//
