.PHONY: dev

dev:
	nodemon --exec go run cmd/main.go --signal SIGTERM
