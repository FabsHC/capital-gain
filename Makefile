build:
	go build cmd/app/main.go

test:
	go test -v ./... -timeout 5m

test-docker:
	docker run -it --rm -v $$PWD:$$PWD -w $$PWD -v /var/run/docker.sock:/var/run/docker.sock golang:1.22 go test ./... -v

lint:
	golangci-lint run --fix

cover:
	gotestsum --format pkgname ./... -coverprofile=coverage.out -coverpkg=./... && \
	grep -v mock coverage.out > tmpcoverage && mv tmpcoverage coverage.out && \
	go tool cover -html=coverage.out -o coverage.html