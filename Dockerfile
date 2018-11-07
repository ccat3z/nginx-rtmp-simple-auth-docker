FROM golang:stretch as authbinary
WORKDIR /src
COPY *.go .
RUN go build -o auth

FROM tiangolo/nginx-rtmp:latest
COPY --from=authbinary /src/auth /usr/bin/auth-server
COPY nginx.conf /etc/nginx/nginx.conf
COPY run.sh /run.sh
EXPOSE 1935
CMD /run.sh
