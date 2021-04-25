FROM golang:latest

RUN mkdir /build

WORKDIR /build

RUN export GO111MODULE=on go get -u goa.design/goa/v3@v3
RUN go get github.com/saiduddukuri411/practice/main
RUN cd /build && git clone https://github.com/saiduddukuri411/practice.git
# https://github.com/AdminTurnedDevOps/GOWebAPI/main

# https://github.com/saiduddukuri411/practice

RUN cd /build/GCP/main && go build

EXPOSE 8080

ENTRYPOINT [ "/build/GCP/main/main" ]