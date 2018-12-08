FROM golang:latest

RUN mkdir -p ${GOPATH}/src/github.com/devdinu
COPY . ${GOPATH}/src/github.com/devdinu/slot_machine

WORKDIR ${GOPATH}/src/github.com/devdinu/slot_machine

RUN env
RUN make install-deps build

EXPOSE 8080
CMD ["./cmd/server/server"]
