FROM golang:1.17.0-alpine3.14
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod tidy -go=1.16 && go mod tidy -go=1.17
RUN go build -o bin/main main.go
EXPOSE 4000
CMD [ "bin/main" ]