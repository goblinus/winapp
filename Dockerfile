FROM golang:alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /task_manager ./cmd/main.go
EXPOSE 8085
CMD [ "/task_manager" ]
