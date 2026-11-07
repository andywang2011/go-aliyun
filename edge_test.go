package aliyun

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewClient_InvalidConfig(t *testing.T) {
	if _, err := NewClient(Config{}); err == nil {
		t.Fatal("expected validate error")
	}
}

func TestNewGreenAPI_Valid(t *testing.T) {
	api, err := NewGreenAPI(testConfig())
	if err != nil {
		t.Fatalf("NewGreenAPI: %v", err)
	}
	if api == nil {
		t.Fatal("expected API")
	}
}

func TestDoRequest_MarshalError(t *testing.T) {
	old := jsonMarshalRequest
	defer func() { jsonMarshalRequest = old }()
	jsonMarshalRequest = func(any) ([]byte, error) { return nil, errors.New("marshal fail") }

	c := testClient(t)
	_, err := c.CheckText(context.Background(), "x")
	if err == nil || !strings.Contains(err.Error(), "marshal") {
		t.Fatalf("err=%v", err)
	}
}

func TestDoRequest_ReadBodyError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Length", "1")
	}))
	defer srv.Close()

	c := testClient(t)
	c.cfg.Host = srv.URL
	_, err := c.CheckText(context.Background(), "x")
	if err == nil {
		t.Fatal("expected read error")
	}
}

func TestDoRequest_NewRequestError(t *testing.T) {
	c := testClient(t)
	_, err := c.doRequest(context.Background(), "/green\x00/text/scan", ScanRequest{
		Scenes: []string{"antispam"},
		Tasks:  []TextTask{{DataId: "1", Content: "x"}},
	})
	if err == nil {
		t.Fatal("expected new request error")
	}
}

func TestGreenError_ErrorWithoutHTTPStatus(t *testing.T) {
	e := &GreenError{Code: 1, Msg: "m"}
	if e.Error() == "" {
		t.Fatal("expected message")
	}
}
