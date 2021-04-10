FROM golang:alpine
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
EXPOSE 8080
RUN go build -o main . 