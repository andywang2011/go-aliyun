package aliyun

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func ExampleClient_CheckText() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/green/text/scan" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		_, _ = w.Write([]byte(`{"code":200,"msg":"OK","data":[]}`))
	}))
	defer srv.Close()

	cfg := testConfig()
	cfg.Host = srv.URL
	api, err := NewClient(cfg)
	if err != nil {
		panic(err)
	}

	resp, err := api.CheckText(context.Background(), "hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Code)
	// Output: 200
}
