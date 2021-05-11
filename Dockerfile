FROM golang:1.16 as packager
WORKDIR /go/src/github.com/sgreaves1/Agent-server
COPY . .
RUN go install

FROM gcr.io/distroless/base
COPY --from=packager /go/bin/Agent-server /

EXPOSE 7621
USER nobody

ENTRYPOINT ["/Agent-server"]