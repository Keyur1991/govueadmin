FROM govueadmin_base:1.0
RUN mkdir -p /go/src/govueadmin/microservices/users
COPY ./microservices/users/ /go/src/govueadmin/microservices/users/
RUN go mod tidy -compat=1.17
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD ["./build.sh"]