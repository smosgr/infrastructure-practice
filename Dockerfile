FROM scratch

ADD ./bin/web-server-dist /web-server

ENTRYPOINT ["/web-server"]
