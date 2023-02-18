FROM govueadmin_base:1.1
RUN mkdir -p /go/src/govueadmin/microservices/roles
COPY ./microservices/roles/ /go/src/govueadmin/microservices/roles/
RUN go mod tidy -compat=1.17