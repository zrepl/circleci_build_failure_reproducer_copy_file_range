FROM gcc:latest

ADD copyfilerange.c /src/copyfilerange.c

RUN gcc -o /copyfilerange_demo /src/copyfilerange.c
RUN echo "foo" > /src.txt && /copyfilerange_demo /src.txt /dst.txt

