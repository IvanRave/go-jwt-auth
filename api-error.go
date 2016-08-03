package main

// apiError for clients
type apiError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// Error satisfy the error interface
func (ae apiError) Error() string {
	return ae.Code
}
