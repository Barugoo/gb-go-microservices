FROM golang:alpine
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
EXPOSE 9094
RUN go build -o main . 
CMD ["/app/main"]