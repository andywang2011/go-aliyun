# go-aliyun

Alibaba Cloud Green content moderation client for Go.

## Install

```bash
go get github.com/andywang2011/go-aliyun@v0.1.0
```

## Quick start

```go
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
```

## Examples

```bash
cd examples
go run .
```

## License

MIT — see [LICENSE](LICENSE).
