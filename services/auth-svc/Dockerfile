FROM golang:1.19-buster as builder

ARG GITHUB_TOKEN

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN git config --global url."https://$GITHUB_TOKEN@github.com/".insteadOf "https://github.com/"
RUN GOOS=linux go build -o /bin/goapp


FROM debian:buster-slim
RUN apt-get update -y \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
        openssl \
        bash \
        curl \
        wget \
        tar \
        gzip \
    && update-ca-certificates \
    && apt-get clean \
    && rm -rf /tmp/* /var/tmp/* /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /bin/goapp goapp-output

EXPOSE 80
EXPOSE 9090
EXPOSE 9800

ENTRYPOINT ["./goapp-output"]