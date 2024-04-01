FROM golang:1.21.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o treasury .

FROM scratch

COPY --from=builder /app/treasury .

EXPOSE 8080

CMD ["/treasury"]
