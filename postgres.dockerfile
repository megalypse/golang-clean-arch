FROM postgres:15.2-alpine

LABEL author="Bruno"
LABEL description="A database for this demo app"
LABEL version="1.0"

COPY db/*.sql /docker-entrypoint-initdb.d/
