FROM scratch

ADD geo-docker-go .

CMD ["./geo-docker-go"]