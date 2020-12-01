FROM golang:1.14.12-stretch as builder
WORKDIR $GOPATH/src/github.com/keremavci/go-helm-cloudbuild
ADD . .
RUN go mod download && \
    go test -v ./... && \
    CGO_ENABLED=0 GOOC=linux GOARCH=amd64 go build -o go-helm-cloudbuild .


FROM scratch
EXPOSE 8080
COPY --from=builder /go/src/github.com/keremavci/go-helm-cloudbuild/go-helm-cloudbuild /go-helm-cloudbuild
ENTRYPOINT ["/go-helm-cloudbuild"]