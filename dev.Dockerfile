FROM workfoxes/calypso:latest
WORKDIR /go/src/github.com/workfoxes/kayo

COPY . /go/src/github.com/workfoxes/kayo

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go bunetworksild kayo.go" --command=./main