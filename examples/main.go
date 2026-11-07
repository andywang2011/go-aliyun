package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	aliyun "github.com/andywang2011/go-aliyun"
)

func main() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/green/text/scan" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		_, _ = w.Write([]byte(`{"code":200,"msg":"OK","data":[]}`))
	}))
	defer srv.Close()

	cfg := aliyun.DefaultConfig()
	cfg.AccessKeyID = "demo-key-id"
	cfg.AccessKeySecret = "demo-key-secret"
	cfg.Host = srv.URL

	api, err := aliyun.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := api.CheckText(context.Background(), "sample text")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("code:", resp.Code)
}
