package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAuthToken(t *testing.T) {
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"token":"test-token"}`))
	}))
	defer mock.Close()

	header := `Bearer realm="` + mock.URL + `",service="registry.docker.io",scope="repository:library/redis:pull"`
	token, err := GetAuthToken(header)
	if err != nil || token != "test-token" {
		t.Fatalf("Expected token 'test-token', got %v (err: %v)", token, err)
	}
}
