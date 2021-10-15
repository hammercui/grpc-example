.PHONY: all build run gotool clean help

proto:
	cd api && sh gen.sh

ser:
	cd server && go run main.go

cli:
	cd client && go run main.go