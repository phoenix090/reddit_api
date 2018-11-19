FROM golang:1.9.2-stretch

RUN adduser --disabled-password --gecos '' api
USER api

WORKDIR $GOPATH/src/reddit_api
COPY . .

RUN go get github.com/gravityblast/fresh
RUN go-wrapper download
RUN go-wrapper install
RUN source private.env

CMD ["fresh"]
