package aliyun

import "testing"

func TestConfig_Validate(t *testing.T) {
	if err := testConfig().Validate(); err != nil {
		t.Fatal(err)
	}
	cases := []Config{
		{AccessKeyID: "id", AccessKeySecret: "sec"},
		{Host: DefaultGreenHost, AccessKeySecret: "sec"},
		{Host: DefaultGreenHost, AccessKeyID: "id"},
	}
	for _, c := range cases {
		if err := c.Validate(); err == nil {
			t.Fatalf("expected error for %+v", c)
		}
	}
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.Host == "" || cfg.Timeout <= 0 {
		t.Fatal(cfg)
	}
}

func TestNewGreenAPI_InvalidConfig(t *testing.T) {
	api, err := NewGreenAPI(Config{})
	if err == nil {
		t.Fatal("expected validation error for invalid config")
	}
	if api != nil {
		t.Fatalf("expected nil API on error, got %v", api)
	}
}
