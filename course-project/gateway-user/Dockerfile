FROM golang:alpine
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
EXPOSE 8080
RUN go install ./cmd/gateway-user-server