FROM golang:1.14 AS builder

WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -o mstest .

FROM scratch
WORKDIR /app
EXPOSE 8080
COPY --from=builder /app/mstest .

CMD ["./mstest"]