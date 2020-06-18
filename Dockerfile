FROM golang:latest
LABEL authors="sarmerer, sarmai" \
    maintainer="sarmerer, sarmai" \
    version="1.0" \
    description="ascii-asrt-web"
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN go build -o main .
EXPOSE 4241
CMD ["./main"]