FROM golang:latest

ADD copyfilerange.go /src/copyfilerange.go


RUN echo "foo" > /src.txt && go run /src/copyfilerange.go /src.txt /dst.txt
RUN echo "foo" > /src.txt && go run /src/copyfilerange.go /src.txt /tmp/dst.txt


