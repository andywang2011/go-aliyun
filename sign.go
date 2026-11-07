package aliyun

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"sort"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	greenMethod         = "POST"
	greenNewline        = "\n"
	greenMIME           = "application/json"
	acsSignatureVersion = "1.0"
	acsSignatureMethod  = "HMAC-SHA1"
	acsVersion          = "2018-05-09"
)

func buildRequestHeader(requestBody, path, accessKeyID, accessKeySecret string) map[string]string {
	md5Str := contentMD5([]byte(requestBody))
	gmtCreate := time.Now().UTC().Format(http.TimeFormat)

	acsHeader := map[string]string{
		"x-acs-signature-method":  acsSignatureMethod,
		"x-acs-signature-nonce":   uuid.NewV4().String(),
		"x-acs-signature-version": acsSignatureVersion,
		"x-acs-version":           acsVersion,
	}

	b := bytes.Buffer{}
	b.WriteString(greenMethod)
	b.WriteString(greenNewline)
	b.WriteString(greenMIME)
	b.WriteString(greenNewline)
	b.WriteString(md5Str)
	b.WriteString(greenNewline)
	b.WriteString(greenMIME)
	b.WriteString(greenNewline)
	b.WriteString(gmtCreate)
	b.WriteString(greenNewline)

	acsKeys := make([]string, 0, len(acsHeader))
	for k := range acsHeader {
		acsKeys = append(acsKeys, k)
	}
	sort.Strings(acsKeys)

	for _, v := range acsKeys {
		b.WriteString(v)
		b.WriteString(":")
		b.WriteString(acsHeader[v])
		b.WriteString(greenNewline)
	}
	b.WriteString(path)

	mac := hmac.New(sha1.New, []byte(accessKeySecret))
	_, _ = mac.Write([]byte(b.String()))
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	authorization := "acs " + accessKeyID + ":" + sign

	headers := map[string]string{
		"Content-Type":  greenMIME,
		"Accept":        greenMIME,
		"Date":          gmtCreate,
		"Content-MD5":   md5Str,
		"Authorization": authorization,
	}
	for k, v := range acsHeader {
		headers[k] = v
	}
	return headers
}

func contentMD5(s []byte) string {
	sum := md5.Sum(s)
	return base64.StdEncoding.EncodeToString(sum[:])
}
