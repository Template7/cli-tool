FROM golang:alpine as builder
ARG PROJECT=cli-tool

WORKDIR /$PROJECT
COPY . .
RUN go mod download
RUN go build -o $PROJECT .

FROM alpine
ARG PROJECT=cli-tool

WORKDIR /$PROJECT
COPY --from=builder /$PROJECT/$PROJECT ./

CMD ["./$PROJECT"]
