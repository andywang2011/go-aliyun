package aliyun

type testingT interface {
	Helper()
	Fatal(...any)
}

func testConfig() Config {
	cfg := DefaultConfig()
	cfg.AccessKeyID = "test-access-key-id"
	cfg.AccessKeySecret = "test-access-key-secret"
	return cfg
}

func testClient(t testingT) *Client {
	t.Helper()
	api, err := NewClient(testConfig())
	if err != nil {
		t.Fatal(err)
	}
	c, ok := api.(*Client)
	if !ok {
		t.Fatal("expected *Client")
	}
	return c
}
