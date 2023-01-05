FROM golang:1.19-buster

WORKDIR /

COPY ./combo /combo
COPY ./resources /resources

EXPOSE 80

# USER nonroot:nonroot

ENTRYPOINT ["/combo"]
