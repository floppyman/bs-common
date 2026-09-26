package httpres

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUnauthorized(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	Unauthorized(rec, "bad creds")

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("status: got %d, want %d", rec.Code, http.StatusUnauthorized)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "bad creds") {
		t.Errorf("body missing message: %s", body)
	}
}

func TestForbidden(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	Forbidden(rec, "no access")

	if rec.Code != http.StatusForbidden {
		t.Errorf("status: got %d, want %d", rec.Code, http.StatusForbidden)
	}
}

func TestNotAcceptable(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	err := NotAcceptable(rec, []string{"GET", "POST"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("status: got %d, want %d", rec.Code, http.StatusMethodNotAllowed)
	}
	allow := rec.Header().Get("Allow")
	if !strings.Contains(allow, "GET") || !strings.Contains(allow, "POST") {
		t.Errorf("Allow header: got %q", allow)
	}
}

func TestBadRequestJson(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	payload := map[string]string{"error": "bad request"}
	err := BadRequestJson(rec, payload)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status: got %d, want %d", rec.Code, http.StatusBadRequest)
	}
	ct := rec.Header().Get("Content-Type")
	if !strings.Contains(ct, "application/json") {
		t.Errorf("content-type: got %q", ct)
	}
	var decoded map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&decoded); err != nil {
		t.Fatalf("failed to decode body: %v", err)
	}
	if decoded["error"] != "bad request" {
		t.Errorf("body: got %v", decoded)
	}
}

func TestNotFoundJson(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	payload := map[string]string{"error": "not found"}
	err := NotFoundJson(rec, payload)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rec.Code != http.StatusNotFound {
		t.Errorf("status: got %d, want %d", rec.Code, http.StatusNotFound)
	}
}

func TestBadRequestString(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	err := BadRequestString(rec, "plain text error")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status: got %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if rec.Body.String() != "plain text error" {
		t.Errorf("body: got %q", rec.Body.String())
	}
}

func TestOkJson(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	payload := map[string]string{"status": "ok"}
	err := OkJson(rec, payload)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("status: got %d, want %d", rec.Code, http.StatusOK)
	}
	ct := rec.Header().Get("Content-Type")
	if !strings.Contains(ct, "application/json") {
		t.Errorf("content-type: got %q", ct)
	}
}

func TestOkString(t *testing.T) {
	t.Parallel()
	rec := httptest.NewRecorder()
	err := OkString(rec, "success")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("status: got %d, want %d", rec.Code, http.StatusOK)
	}
	if rec.Body.String() != "success" {
		t.Errorf("body: got %q", rec.Body.String())
	}
}

func TestResponseContentTypesOverHTTP(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name        string
		handler     http.HandlerFunc
		wantContain string
	}{
		{
			name: "OkJson sets application/json",
			handler: func(w http.ResponseWriter, r *http.Request) {
				_ = OkJson(w, map[string]string{"status": "ok"})
			},
			wantContain: "application/json",
		},
		{
			name: "BadRequestJson sets application/json",
			handler: func(w http.ResponseWriter, r *http.Request) {
				_ = BadRequestJson(w, map[string]string{"error": "bad"})
			},
			wantContain: "application/json",
		},
		{
			name: "NotFoundJson sets application/json",
			handler: func(w http.ResponseWriter, r *http.Request) {
				_ = NotFoundJson(w, map[string]string{"error": "missing"})
			},
			wantContain: "application/json",
		},
		{
			name: "BadRequestString sets text/plain",
			handler: func(w http.ResponseWriter, r *http.Request) {
				_ = BadRequestString(w, "oops")
			},
			wantContain: "text/plain",
		},
		{
			name: "OkString sets text/plain",
			handler: func(w http.ResponseWriter, r *http.Request) {
				_ = OkString(w, "fine")
			},
			wantContain: "text/plain",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(tc.handler)
			defer srv.Close()
			resp, err := http.Get(srv.URL)
			if err != nil {
				t.Fatalf("GET failed: %v", err)
			}
			defer resp.Body.Close()
			ct := resp.Header.Get("Content-Type")
			if !strings.Contains(ct, tc.wantContain) {
				t.Errorf("Content-Type got %q, want substring %q", ct, tc.wantContain)
			}
		})
	}
}

func TestNotAcceptableAllowOverHTTP(t *testing.T) {
	t.Parallel()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = NotAcceptable(w, []string{"GET", "POST"})
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()

	req, err := http.NewRequest(http.MethodGet, srv.URL, nil)
	if err != nil {
		t.Fatalf("build request: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("GET failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("status: got %d, want %d", resp.StatusCode, http.StatusMethodNotAllowed)
	}
	allow := resp.Header.Get("Allow")
	if !strings.Contains(allow, "GET") || !strings.Contains(allow, "POST") {
		t.Errorf("Allow got %q, expected to contain GET and POST", allow)
	}
}
