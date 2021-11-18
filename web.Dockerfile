FROM workfoxes/calypso

WORKDIR /go/src/github.com/workfoxes/kayo

COPY . .

RUN go mod download
RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8000

ENTRYPOINT [ "go", "run", "cmd/server/main.go" ]