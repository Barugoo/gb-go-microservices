FROM golang:alpine
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
EXPOSE 8089
RUN go build -o main . 
CMD ["/app/main"]