package aliyun

import (
	"fmt"
	"time"
)

const (
	DefaultGreenHost = "https://green.cn-shanghai.aliyuncs.com"
	GreenSuccessCode = 200
)

// Config holds Alibaba Cloud Green (content moderation) client settings.
// Fill this struct in the application layer (not from viper inside this library).
type Config struct {
	// Host is the Green API endpoint (region-specific).
	Host string
	// AccessKeyID is the Alibaba Cloud access key ID.
	AccessKeyID string
	// AccessKeySecret is the Alibaba Cloud access key secret.
	AccessKeySecret string
	// Timeout is the HTTP client timeout.
	Timeout time.Duration
}

// DefaultConfig returns defaults for tests and local development.
func DefaultConfig() Config {
	return Config{
		Host:    DefaultGreenHost,
		Timeout: 10 * time.Second,
	}
}

// ConfigDefault is an alias for DefaultConfig (naming consistency across go-* libs).
func ConfigDefault() Config { return DefaultConfig() }

// Validate checks required fields.
func (c Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("aliyun green config: host is required")
	}
	if c.AccessKeyID == "" {
		return fmt.Errorf("aliyun green config: access_key_id is required")
	}
	if c.AccessKeySecret == "" {
		return fmt.Errorf("aliyun green config: access_key_secret is required")
	}
	return nil
}
