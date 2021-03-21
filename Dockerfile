ARG alpine_version
FROM arm64v8/alpine:latest
RUN apk add --update 
RUN apk add --no-cache bluez bluez-deprecated alsa-utils alsa-utils-doc alsa-lib alsaconf
#RUN apk add --no-cache -X http://dl-cdn.alpinelinux.org/alpine/edge/testing bluez-alsa
RUN apk add --no-cache go
RUN apk add git libc-dev go --repository http://dl-cdn.alpinelinux.org/alpine/v3.10/main/ --repository http://dl-cdn.alpinelinux.org/alpine/v3.10/community/
RUN go get github.com/paypal/gatt


