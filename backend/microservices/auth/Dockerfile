FROM govueadmin_base:1.1
RUN mkdir -p /go/src/govueadmin/microservices/auth
COPY ./microservices/auth/ /go/src/govueadmin/microservices/auth/
RUN rm -rf /go/src/govueadmin/microservices/common/client/auth
RUN rm -rf /go/src/govueadmin/microservices/common/pb/auth
RUN rm -rf /go/src/govueadmin/microservices/common/middlewares/auth
RUN go mod tidy -compat=1.17