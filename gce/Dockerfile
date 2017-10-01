FROM golang:1.8-alpine
ADD . /go/src/gce-example
RUN go install gce-example

FROM alpine:latest
COPY --from=0 /go/bin/gce-example .
ENV PORT 8080
CMD ["./gce-example"]
