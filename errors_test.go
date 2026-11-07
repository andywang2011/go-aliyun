package aliyun

import (
	"errors"
	"testing"
)

func TestGreenResp_IsSuccess(t *testing.T) {
	if !(&GreenResp{Code: GreenSuccessCode}).IsSuccess() {
		t.Fatal("expected success")
	}
	if (&GreenResp{Code: 400}).IsSuccess() {
		t.Fatal("expected failure")
	}
	if (*GreenResp)(nil).IsSuccess() {
		t.Fatal("nil resp")
	}
}

func TestGreenError_ErrorAndIs(t *testing.T) {
	err := &GreenError{HTTPStatus: 500, Code: 1, Msg: "fail", RawBody: "raw"}
	if err.Error() == "" {
		t.Fatal("expected message")
	}
	got, ok := IsGreenError(err)
	if !ok || got.HTTPStatus != 500 {
		t.Fatalf("got=%+v ok=%v", got, ok)
	}
	if _, ok := IsGreenError(errors.New("other")); ok {
		t.Fatal("expected false")
	}
}
