FROM golang:1.17
RUN mkdir -p /go/src/govueadmin/microservices
RUN mkdir -p /go/src/govueadmin/microservices/auth
RUN mkdir -p /go/src/govueadmin/framework
COPY ./framework/ /go/src/govueadmin/framework/
COPY ./microservices/auth/ /go/src/govueadmin/microservices/auth/
WORKDIR /go/src/govueadmin
RUN apt-get update
RUN apt install -y protobuf-compiler
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=direct
RUN go mod init
RUN go mod tidy -compat=1.17
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
RUN export PATH="$PATH:$(go env GOPATH)/bin"
CMD ["./build.sh"]