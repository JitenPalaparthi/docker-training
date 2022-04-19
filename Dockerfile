FROM ubuntu

LABEL "author"="Jiten"

LABEL "version"="v0.0.1"

RUN apt-get update -y && \
 apt-get upgrade -y && \
 apt-get install vim -y

RUN echo "Hello World"
