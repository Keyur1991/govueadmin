FROM golang:1.17
RUN mkdir -p /go/src/govueadmin/microservices
RUN mkdir -p /go/src/govueadmin/microservices/users
RUN mkdir -p /go/src/govueadmin/framework
COPY ./framework/ /go/src/govueadmin/framework/
COPY ./microservices/users/ /go/src/govueadmin/microservices/users/
WORKDIR /go/src/govueadmin
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=direct
RUN go mod init
RUN go mod tidy
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD ["./build.sh"]