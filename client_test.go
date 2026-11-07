package aliyun

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewClient_Success(t *testing.T) {
	c := testClient(t)
	if c.cfg.AccessKeyID != "test-access-key-id" {
		t.Fatal(c.cfg)
	}
}

func TestClient_CheckText_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/green/text/scan" {
			t.Fatalf("path=%s", r.URL.Path)
		}
		if r.Header.Get("Authorization") == "" {
			t.Fatal("missing authorization")
		}
		_, _ = w.Write([]byte(`{"code":200,"msg":"OK","data":[]}`))
	}))
	defer srv.Close()

	cfg := testConfig()
	cfg.Host = srv.URL
	api, err := NewClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := api.CheckText(context.Background(), "hello")
	if err != nil || !resp.IsSuccess() {
		t.Fatalf("resp=%+v err=%v", resp, err)
	}
}

func TestClient_CheckImage_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/green/image/scan" {
			t.Fatalf("path=%s", r.URL.Path)
		}
		_, _ = w.Write([]byte(`{"code":200,"msg":"OK","data":[]}`))
	}))
	defer srv.Close()

	cfg := testConfig()
	cfg.Host = srv.URL
	api, err := NewClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := api.CheckImage(context.Background(), "https://example.com/a.png")
	if err != nil || !resp.IsSuccess() {
		t.Fatalf("resp=%+v err=%v", resp, err)
	}
}

func TestClient_doRequest_HTTPError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("upstream error"))
	}))
	defer srv.Close()

	c := testClient(t)
	c.cfg.Host = srv.URL
	_, err := c.CheckText(context.Background(), "x")
	if err == nil {
		t.Fatal("expected error")
	}
	ge, ok := IsGreenError(err)
	if !ok || ge.HTTPStatus != http.StatusInternalServerError {
		t.Fatalf("err=%v ge=%+v", err, ge)
	}
}

func TestClient_doRequest_InvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("not-json"))
	}))
	defer srv.Close()

	c := testClient(t)
	c.cfg.Host = srv.URL
	_, err := c.CheckText(context.Background(), "x")
	if err == nil || !strings.Contains(err.Error(), "unmarshal") {
		t.Fatalf("err=%v", err)
	}
}

func TestClient_doRequest_InvalidGateway(t *testing.T) {
	c := testClient(t)
	c.cfg.Host = "://bad-url"
	_, err := c.CheckText(context.Background(), "x")
	if err == nil {
		t.Fatal("expected request error")
	}
}

func TestNewClient_DefaultTimeout(t *testing.T) {
	cfg := Config{
		Host:            DefaultGreenHost,
		AccessKeyID:     testConfig().AccessKeyID,
		AccessKeySecret: testConfig().AccessKeySecret,
	}
	api, err := NewClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if api.(*Client).cfg.Timeout != 10*time.Second {
		t.Fatalf("timeout=%v", api.(*Client).cfg.Timeout)
	}
}
