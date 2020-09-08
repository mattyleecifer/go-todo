# FROM golang:latest
# FROM golang:1.15
FROM golang

WORKDIR /go/src/app

RUN git clone https://github.com/mattyleecifer/go-todo.git /go/src/app

RUN rm README.md

RUN go get -d -v ./...

RUN go install -v ./...

CMD app


# build it
# docker build -t scraper .
# docker build -t --no-cache scraper .

# run it interactive
# docker run -it scraper

# run it interactive and keep files
# docker run scraper -v "$PWD":/go/src/app
