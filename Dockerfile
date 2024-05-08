FROM golang:1.22 

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-market-warehouse-api ./cmd/main.go

EXPOSE 8000

CMD [ "/go-market-warehouse-api" ]
