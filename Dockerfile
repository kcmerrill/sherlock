FROM golang:1.6
MAINTAINER kc merrill <kcmerrill@gmail.com>
COPY . /go/src/github.com/kcmerrill/sherlock
WORKDIR /go/src/github.com/kcmerrill/sherlock
RUN  go build -ldflags "-X main.Commit=`git rev-parse HEAD` -X main.Version=0.1.`git rev-list --count HEAD`" -o /usr/local/bin/alfred
EXPOSE 80 
ENTRYPOINT ["sherlock"]
