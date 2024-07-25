FROM golang:1.22.5

WORKDIR app/

COPY .. .

RUN go mod download && go mod verify

RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

CMD ["sh", "-c", "go run cmd/main/main.go"]