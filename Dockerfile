FROM golang:alpine
WORKDIR /app
COPY * ./
RUN go mod tidy
RUN go build -o /winapp
EXPOSE 8085
CMD [ "/winapp" ]
