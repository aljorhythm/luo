setup:
	# githooks
	git config core.hooksPath .githooks
	go mod tidy

lint:
	go fmt ./...

microtest:
	go test ./... -v