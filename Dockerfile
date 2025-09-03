FROM surnet/alpine-wkhtmltopdf:3.20.2-0.12.6-full as builder

FROM golang:1.23.7-alpine3.20

RUN mkdir /app
RUN apk update && \
    apk add --no-cache git openssh tzdata build-base python3 net-tools

RUN  apk update && apk add --no-cache \
      libstdc++ \
      libx11 \
      libxrender \
      libxext \
      libressl \
      ca-certificates \
      fontconfig \
      freetype \
      ttf-dejavu \
      ttf-droid \
      ttf-freefont \
      ttf-liberation \
    && apk add --no-cache --virtual .build-deps \
      msttcorefonts-installer \
    \
    # Install Microsoft fonts
    && update-ms-fonts \
    && fc-cache -f \
    \
    # Clean up when done
    && rm -rf /var/cache/apk/* \
    && rm -rf /tmp/* \
    && apk del .build-deps

WORKDIR /app

COPY .env.example .env
COPY . .

RUN go install github.com/buu700/gin@latest
RUN GO111MODULE=auto
RUN go mod tidy

RUN make build

ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY --from=builder /bin/wkhtmltopdf /bin/wkhtmltopdf
COPY --from=builder /bin/wkhtmltoimage /bin/wkhtmltoimage

WORKDIR /app
EXPOSE 8003

ENTRYPOINT ["/app/payment-service"]