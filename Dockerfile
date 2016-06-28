#FROM golang
#
#COPY . /go/src/
#
#RUN go get -d -v
#RUN go install -v
#
#ENTRYPOINT /go/bin/cimon_jobs
#
#EXPOSE 10000

FROM golang

#COPY ../../.. go/src

ADD . /go/src/github.com/andreasmaier/cimon_jobs

RUN go get -u github.com/golang/protobuf/proto
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/gengo/grpc-gateway/runtime
RUN go get -u github.com/philips/grpc-gateway-example/insecure
RUN go get -u github.com/ziutek/mymysql/godrv

RUN go install github.com/andreasmaier/cimon_jobs

#RUN go install github.com/andreasmaier/cimon_jobs

ENTRYPOINT /go/bin/cimon_jobs

EXPOSE 10000