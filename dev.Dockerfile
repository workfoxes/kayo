FROM golang:latest
COPY . /kayo
COPY go-dev.mod /kayo/go.mod
COPY ../gobase /gobase
WORKDIR /kayo

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go bunetworksild kayo.go" --command=./main