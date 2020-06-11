FROM golang:1.10.3 AS builder

LABEL authors="sarmerer, sarmai" \
    maintainer="sarmerer, sarmai" \
    version="1.0" \
    description="ascii-asrt-web"

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./ascii-art-web"]