FROM alpine:latest

RUN apk add --update ca-certificates
RUN update-ca-certificates
RUN apk add --update tzdata
ENV TZ=Asia/Shanghai
COPY client /
RUN  mkdir /public
RUN chmod +x /client

CMD ["/client"]
