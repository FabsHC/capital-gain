test-go:
	@go test -v ./... -timeout 5m
test-docker:
	docker run -it --rm -v $$PWD:$$PWD -w $$PWD -v /var/run/docker.sock:/var/run/docker.sock golang:1.20.3 go test ./... -v