FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /marketplace-app ./cmd/main.go

FROM gcr.io/distroless/static-debian12

COPY --from=builder /marketplace-app /marketplace-app

CMD ["/marketplace-app"]