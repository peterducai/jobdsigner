#!/bin/bash

go build -o jobdsigner jobdsigner.go
go install .
rm -rf jobdsigner
sleep 1
jobdsigner

	# rm -rf ~/go/src/github.com/peterducai/jobdsigner
	# go get github.com/peterducai/jobdsigner

# compile:
# 	# 32-Bit Systems
# 	# FreeBDS
# 	GOOS=freebsd GOARCH=386 go build -o bin/jobdsigner-freebsd-386 jobdsigner.go
# 	# MacOS
# 	GOOS=darwin GOARCH=386 go build -o bin/jobdsigner-darwin-386 jobdsigner.go
# 	# Linux
# 	GOOS=linux GOARCH=386 go build -o bin/jobdsigner-linux-386 jobdsigner.go
# 	# Windows
# 	GOOS=windows GOARCH=386 go build -o bin/jobdsigner-windows-386 jobdsigner.go
#         # 64-Bit
# 	# FreeBDS
# 	GOOS=freebsd GOARCH=amd64 go build -o bin/jobdsigner-freebsd-amd64 jobdsigner.go
# 	# MacOS
# 	GOOS=darwin GOARCH=amd64 go build -o bin/jobdsigner-darwin-amd64 jobdsigner.go
# 	# Linux
# 	GOOS=linux GOARCH=amd64 go build -o bin/jobdsigner-linux-amd64 jobdsigner.go
# 	# Windows
# 	GOOS=windows GOARCH=amd64 go build -o bin/jobdsigner-windows-amd64 jobdsigner.go
