package domains

// Employee category type
type Employee struct {
	ID         string     `json:"id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"name" `
	Role       string     `json:"role" `
	Email      string     `json:"email" `
	Department string     `json:"department" `
}
