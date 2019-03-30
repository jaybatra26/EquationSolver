FROM golang:1.7

RUN apt-get -y update
RUN mkdir -p /equationsolver
WORKDIR equationsolver/
ADD ./main .
CMD ["./main"]
