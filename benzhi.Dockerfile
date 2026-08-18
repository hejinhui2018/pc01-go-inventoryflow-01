FROM golang:1.23.12-bookworm
WORKDIR /src
COPY go.mod ./
COPY . .
RUN go test ./... -count=1
RUN go vet ./...
RUN go build ./...
CMD ["go", "run", "./cmd/inventoryflow", "-port", "8080"]
