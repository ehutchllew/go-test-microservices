# Build Docker Image
FROM alpine:latest

RUN mkdir /app

COPY authApp /app

CMD ["/app/authApp"]
