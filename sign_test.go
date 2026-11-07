package aliyun

import (
	"strings"
	"testing"
)

func TestContentMD5(t *testing.T) {
	got := contentMD5([]byte(`{"a":1}`))
	if got == "" {
		t.Fatal("expected md5")
	}
}

func TestBuildRequestHeader(t *testing.T) {
	body := `{"scenes":["antispam"]}`
	h := buildRequestHeader(body, "/green/text/scan", "ak", "sk")
	for _, key := range []string{"Content-Type", "Accept", "Date", "Content-MD5", "Authorization", "x-acs-version"} {
		if h[key] == "" {
			t.Fatalf("missing header %s: %+v", key, h)
		}
	}
	if !strings.HasPrefix(h["Authorization"], "acs ak:") {
		t.Fatalf("authorization=%q", h["Authorization"])
	}
}
