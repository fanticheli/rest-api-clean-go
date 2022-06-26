FROM golang:1.16

WORKDIR /home/golang/app

CMD ["tail", "-f", "/dev/null"]
