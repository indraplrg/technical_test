package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// parseID extracts and converts the ":id" path parameter to uint.
func parseID(c *gin.Context) (uint, error) {
	value, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}

// parseUint parses a string to uint, returning an error when invalid.
func parseUint(value string) (uint, error) {
	parsed, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(parsed), nil
}

// parseOrDefault parses an integer query parameter with a fallback.
func parseOrDefault(value string, fallback int) int {
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}
