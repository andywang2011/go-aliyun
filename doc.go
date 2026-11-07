// Package aliyun provides a testable client for Alibaba Cloud Green content moderation.
//
// Design constraints:
//   - Config injection via Config struct (no viper; map keys in the application).
//   - ACS HMAC-SHA1 request signing for JSON POST to the Green endpoint.
//   - CheckText and CheckImage for antispam scanning.
//
// This package is not related to Ant Group Alipay (see go-alipay).
package aliyun
