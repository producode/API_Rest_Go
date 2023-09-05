FROM golang:1.20-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
ENV PORT 30501
ENV PORT_CORS 30501
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
EXPOSE 30501

# Run
CMD ["/docker-gs-ping"]