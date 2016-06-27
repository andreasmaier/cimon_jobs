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

#COPY . /go/src/github.com/andreasmaier/cimon_jobs

RUN go get -u github.com/golang/protobuf/proto
RUN go get -u github.com/golang/protobuf/protoc-gen-go

RUN go get github.com/andreasmaier/cimon_jobs

#RUN go install github.com/andreasmaier/cimon_jobs

ENTRYPOINT /go/bin/cimon_jobs

EXPOSE 10000