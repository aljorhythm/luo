setup:
	# githooks
	git config core.hooksPath .githooks
	go mod tidy

lint:
	sh .format.sh

microtest:
	go test ./... -v