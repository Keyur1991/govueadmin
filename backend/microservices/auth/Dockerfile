FROM app-base:1.0
RUN mkdir -p /go/src/app/microservices/auth
COPY ./microservices/auth/ /go/src/app/microservices/auth/
RUN rm -rf /go/src/app/microservices/common/client/auth
RUN rm -rf /go/src/app/microservices/common/pb/auth
RUN rm -rf /go/src/app/microservices/common/middlewares/auth
WORKDIR /go/src/app
CMD [ "go", "run","./microservices/auth/main.go"]