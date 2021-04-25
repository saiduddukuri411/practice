FROM golang:latest

RUN mkdir /build

WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/AdminTurnedDevOps/GOWebAPI/main
RUN cd /build && git clone https://github.com/AdminTurnedDevOps/GOWebAPI.git

RUN cd /build/GOWebAPI/main && go build

EXPOSE 8080

ENTRYPOINT [ "/build/GCP/main/main" ]