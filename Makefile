run-dev:
	nodemon --exec go run cmd/server.go --signal SIGTERM || exit 1
.PHONY: run-dev