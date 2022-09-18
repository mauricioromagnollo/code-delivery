FROM golang:1.16

RUN useradd -ms /bin/bash go
RUN chown go /go/src
WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN apt-get update && \
  apt-get install build-essential librdkafka-dev -y

USER go

CMD [ "tail", "-f", "/dev/null" ]
