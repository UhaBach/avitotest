FROM golang:latest 

RUN go version 
ENV GOPATH=/
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o go-app ./cmd/apiserver/main.go
CMD ["./go-app"]