FROM busybox:latest
COPY ./server /server
CMD ["/server"]
