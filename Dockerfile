FROM golang:1.17
WORKDIR /go/app
RUN apt update && apt install build-essential librdkafka-dev -y
CMD ["tail","-f","/dev/null"]