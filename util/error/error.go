package error

type (
	// Error ...
	Error struct {
		Status  int    `json:"status,omitempty"`
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}
)

// Error ...
func (e Error) Error() string {
	return e.Message
}
