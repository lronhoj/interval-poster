FROM golang:1.9

# cause dependencies to be statically linked
ENV CGO_ENABLED=0

WORKDIR /go/src/app
COPY . .
RUN go build -o interval-poster main.go


FROM scratch
COPY --from=0 /go/src/app/interval-poster /
CMD ["/interval-poster"]