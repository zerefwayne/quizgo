FROM golang:1.8.5-jessie

RUN mkdir app

WORKDIR /app

ENV SRC_DIR=/go/src/github.com/zerefwayne/quizgo/

ADD . $SRC_DIR

RUN cd $SRC_DIR

COPY . /app

CMD ["go", "run", "main.go"]
