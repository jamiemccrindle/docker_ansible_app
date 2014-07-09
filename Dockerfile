FROM google/golang

WORKDIR /gopath/src/app
ADD . /gopath/src/app/
RUN go get app

EXPOSE 8000

CMD []
ENTRYPOINT ["/gopath/bin/app"]
