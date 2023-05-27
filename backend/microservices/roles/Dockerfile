FROM app-base:1.0
RUN mkdir -p /go/src/app/microservices/roles
COPY ./microservices/roles/ /go/src/app/microservices/roles/
WORKDIR /go/src/app
CMD ["go", "run", "./microservices/roles/main.go"]