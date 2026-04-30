FROM golang:1.22

WORKDIR /app

COPY . .

ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org

RUN go mod tidy
RUN go build -o main ./cmd

CMD ["./main"]