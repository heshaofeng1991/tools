FROM golang:1.18-buster

WORKDIR /

COPY ./ddd-sample /ddd-sample
COPY ./resources /resources

EXPOSE 80

# USER nonroot:nonroot

ENTRYPOINT ["/oms-backend"]
