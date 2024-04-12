FROM alpine:3.10
COPY ./server /server
CMD ["/server"]
