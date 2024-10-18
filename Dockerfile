FROM golang:1.22
WORKDIR /app
COPY . .
RUN go test -v ./...
WORKDIR /app/cmd
RUN go build -o capital-gain .
ENTRYPOINT [ "/app/cmd/capital-gain"]