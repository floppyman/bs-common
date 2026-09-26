package utils

import (
	"net/http"
	"strings"
	"testing"
)

func TestColorMethod(t *testing.T) {
	t.Parallel()

	cases := []struct {
		method string
		want   string
	}{
		{http.MethodGet, "GET"},
		{http.MethodPost, "POST"},
		{http.MethodPut, "PUT"},
		{http.MethodDelete, "DELETE"},
		{http.MethodPatch, "PATCH"},
		{"UNKNOWN", "UNKNOWN"},
	}

	for _, tc := range cases {
		got := ColorMethod(tc.method)
		if !strings.Contains(got, tc.want) {
			t.Errorf("ColorMethod(%q): expected to contain %q, got %q", tc.method, tc.want, got)
		}
		// All outputs should contain ANSI reset
		if !strings.Contains(got, "\033[0m") {
			t.Errorf("ColorMethod(%q): missing ANSI reset", tc.method)
		}
	}
}

func TestColorStatus(t *testing.T) {
	t.Parallel()

	cases := []struct {
		code int
		want string
	}{
		{200, "200"},
		{301, "301"},
		{404, "404"},
		{500, "500"},
		{0, "0"},
	}

	for _, tc := range cases {
		got := ColorStatus(tc.code)
		if !strings.Contains(got, tc.want) {
			t.Errorf("ColorStatus(%d): expected to contain %q, got %q", tc.code, tc.want, got)
		}
		if !strings.Contains(got, "\033[0m") {
			t.Errorf("ColorStatus(%d): missing ANSI reset", tc.code)
		}
	}
}
