FROM golang:1.6-onbuild

ENTRYPOINT /go/bin/cimon_jobs

EXPOSE 10000