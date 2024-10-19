FROM golang:1.22-alpine

ENV CGO_ENABLED=0
RUN apk add curl tar make git
COPY . /app
WORKDIR /app

RUN go mod download
RUN make

ENV PORTFWD_CONFIG_FILE_PATH=/app/config.yaml

CMD ["./portfwd"]

