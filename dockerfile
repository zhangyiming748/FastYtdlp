FROM golang:1.24.6-alpine3.22 as builder
WORKDIR /app
COPY . .
# RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#http://mirrors.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories
RUN go env -w GO111MODULE=on
# RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN go build -o ytd main.go

FROM python:3.12.11-alpine3.22
COPY --from=builder /app/ytd /usr/bin/ytd
RUN apk add ffmpeg mediainfo
RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#http://mirrors.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories
# RUN go env -w GO111MODULE=on
# RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN python -m pip install --upgrade pip
# RUN pip config set global.index-url https://mirrors.tuna.tsinghua.edu.cn/pypi/web/simple
RUN python3 -m pip install -U "yt-dlp[default]"
# ENTRYPOINT [ "/usr/bin/ytd" ]
CMD [ "/usr/bin/ytd" ]
