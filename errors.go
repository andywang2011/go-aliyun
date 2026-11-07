package aliyun

import (
	"errors"
	"fmt"
)

// GreenError represents an HTTP or gateway error from Green.
type GreenError struct {
	HTTPStatus int
	Code       int
	Msg        string
	RawBody    string
}

func (e *GreenError) Error() string {
	if e.HTTPStatus > 0 {
		return fmt.Sprintf("aliyun green http=%d code=%d msg=%s", e.HTTPStatus, e.Code, e.Msg)
	}
	return fmt.Sprintf("aliyun green code=%d msg=%s", e.Code, e.Msg)
}

// IsGreenError reports whether err is a *GreenError.
func IsGreenError(err error) (*GreenError, bool) {
	var ge *GreenError
	if errors.As(err, &ge) {
		return ge, true
	}
	return nil, false
}

// IsSuccess reports whether the Green business response succeeded (code == 200).
func (r *GreenResp) IsSuccess() bool {
	return r != nil && r.Code == GreenSuccessCode
}
