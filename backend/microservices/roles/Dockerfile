FROM golang:1.17
RUN mkdir -p /go/src/govueadmin/microservices
RUN mkdir -p /go/src/govueadmin/microservices/roles
RUN mkdir -p /go/src/govueadmin/framework
COPY ./framework/ /go/src/govueadmin/framework/
COPY ./microservices/roles/ /go/src/govueadmin/microservices/roles/
WORKDIR /go/src/govueadmin
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=direct
RUN go mod init
RUN go mod tidy -compat=1.17
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD ["./build.sh"]