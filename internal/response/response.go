package response

import "github.com/gin-gonic/gin"

// Result is the consistent JSON envelope for all API responses.
type Result struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// Success writes a successful response.
func Success(c *gin.Context, httpStatus int, message string, data interface{}) {
	c.JSON(httpStatus, Result{Success: true, Message: message, Data: data})
}

// Error writes an error response. errors is optional and usually holds
// field validation messages.
func Error(c *gin.Context, httpStatus int, message string, errors ...interface{}) {
	var errs interface{}
	if len(errors) > 0 {
		errs = errors[0]
	}
	c.JSON(httpStatus, Result{Success: false, Message: message, Errors: errs})
}

// Paginated holds the data items together with pagination metadata,
// used by list endpoints.
type Paginated struct {
	Items      interface{} `json:"items"`
	Pagination interface{} `json:"pagination"`
}
