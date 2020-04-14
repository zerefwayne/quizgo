FROM golang:1.8.5-jessie

RUN mkdir app

WORKDIR /app

ADD problems /app/problems

ENV SRC_DIR=/go/src/github.com/zerefwayne/quizgo/

ADD . $SRC_DIR

RUN cd $SRC_DIR; go build -o go-quiz; cp go-quiz /app/

ENTRYPOINT ["./go-quiz"]
