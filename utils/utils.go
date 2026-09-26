package utils

import (
	"fmt"
	"net/http"
)

func ColorMethod(method string) string {
	str := fmt.Sprintf("%-6s", method)
	switch method {
	case http.MethodGet:
		return Green(str)
	case http.MethodPost:
		return Yellow(str)
	case http.MethodPut:
		return Cyan(str)
	case http.MethodDelete:
		return Red(str)
	case http.MethodPatch:
		return Blue(str)
	default:
		return White(str)
	}
}

func ColorStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return Green(ToString(code))
	case code >= 300 && code < 400:
		return Cyan(ToString(code))
	case code >= 400 && code < 500:
		return Yellow(ToString(code))
	case code >= 500:
		return Red(ToString(code))
	default:
		return White(ToString(code))
	}
}
