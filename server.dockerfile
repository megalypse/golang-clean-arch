FROM golang:1.20.2-alpine3.17

COPY ./bin/server /bin/server
ENTRYPOINT [ "/bin/server" ]
