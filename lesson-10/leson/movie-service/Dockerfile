FROM golang:alpine
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
EXPOSE 9098
RUN go build -o main . 