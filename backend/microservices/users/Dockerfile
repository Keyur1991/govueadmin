FROM app-base:1.0
RUN mkdir -p /go/src/app/microservices/users
COPY ./microservices/users/ /go/src/app/microservices/users/
WORKDIR /go/src/app
CMD ["go", "run", "./microservices/users/main.go"]