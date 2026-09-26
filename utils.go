package utils

import (
	"fmt"
	"net/http"
)

func ColorMethod(method string) string {
	var c string
	switch method {
	case http.MethodGet:
		c = "\033[32m"
	case http.MethodPost:
		c = "\033[34m"
	case http.MethodPut:
		c = "\033[33m"
	case http.MethodDelete:
		c = "\033[31m"
	case http.MethodPatch:
		c = "\033[36m"
	default:
		c = "\033[37m"
	}
	return fmt.Sprintf("%s%-6s\033[0m", c, method)
}

func ColorStatus(code int) string {
	var c string
	switch {
	case code >= 200 && code < 300:
		c = "\033[32m"
	case code >= 300 && code < 400:
		c = "\033[36m"
	case code >= 400 && code < 500:
		c = "\033[33m"
	case code >= 500:
		c = "\033[31m"
	default:
		c = "\033[37m"
	}
	return fmt.Sprintf("%s%d\033[0m", c, code)
}
