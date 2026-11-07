module github.com/andywang2011/go-aliyun/examples

go 1.27.3

require github.com/andywang2011/go-aliyun v0.0.0

require (
	github.com/andywang2011/go-httplib v0.1.0 // indirect
	github.com/andywang2011/go-logging v0.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.16.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
)

replace (
	github.com/andywang2011/go-httplib => ../../go-httplib
	github.com/andywang2011/go-logging => ../../go-logging
)

replace github.com/andywang2011/go-aliyun => ..
