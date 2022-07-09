# Build stage
FROM golang:1.17-alpine3.13 As builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-386.tar.gz | tar xvz 
          

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY  --from=builder /app/main .
COPY --from=builder  /app/migrate ./migrate
COPY app.env . 
COPY start.sh . 
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]

