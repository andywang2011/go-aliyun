package aliyun

import "context"

// API is the Alibaba Cloud Green (content moderation) client interface.
type API interface {
	CheckText(ctx context.Context, content string) (*GreenResp, error)
	CheckImage(ctx context.Context, imageURL string) (*GreenResp, error)
}
