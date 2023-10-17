FROM golang:1.18-alpine as builder
MAINTAINER Jefri Herdi Triyanto, jefriherditriyanto@gmail.com

#-> Setup Environment
# ENV GOPATH /go
# ENV PATH $PATH:$GOPATH/bin
ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV CGO 0

#-> 🌊 Install Require
RUN apk add --no-cache \
    gcc \
    musl-dev

WORKDIR /build
COPY . .

#-> 🌊 Install Golang Module
RUN go mod download

#-> ⚒️ Build App
RUN go build -o ./run

#-> 💯 Configuration Environment (change target env)
RUN sed -i 's/localhost/host.docker.internal/g' .env

# 🚀 Finishing !!
FROM alpine:latest as runner
WORKDIR /app

# Add the community repository to get ffmpeg
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories

# Install ffmpeg along with the other tools
RUN apk add --no-cache openssl curl nano ffmpeg

COPY --from=builder /build/run  /app/run

RUN chmod +x ./run

ENTRYPOINT ["/app/run"]
CMD ["run"]
