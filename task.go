package aliyun

const (
	// LabelNormal is the moderation label for clean content.
	LabelNormal = "normal"
)

// TextTask is one text moderation task.
type TextTask struct {
	DataId  string `json:"dataId"`  // unique data ID
	Content string `json:"content"` // text to scan
}

// ImageTask is one image moderation task.
type ImageTask struct {
	DataId string `json:"dataId"` // unique data ID
	Url    string `json:"url"`    // image URL
}

// ScanRequest is the Green content moderation request body.
type ScanRequest struct {
	Scenes []string    `json:"scenes"` // scenes to run, e.g. antispam
	Tasks  interface{} `json:"tasks"`  // []TextTask or []ImageTask
}

// GreenResp is the Green scan API response.
type GreenResp struct {
	Code int    `json:"code"` // 200 means success
	Msg  string `json:"msg"`
	Data []struct {
		Results []struct {
			Details []struct {
				Contexts []struct {
					Context   string `json:"context"`
					LibCode   string `json:"libCode"`
					LibName   string `json:"libName"`
					Positions []struct {
						EndPos   int `json:"endPos"`
						StartPos int `json:"startPos"`
					} `json:"positions"`
				} `json:"contexts"`
				Label string `json:"label"`
			} `json:"details"`
			Label      string  `json:"label"`
			Rate       float64 `json:"rate"`
			Scene      string  `json:"scene"`
			Suggestion string  `json:"suggestion"`
		} `json:"results"`
	} `json:"data"`
}

// Request is an alias for ScanRequest for backward compatibility.
type Request = ScanRequest
