FROM golang:1.16-alpine AS builder

WORKDIR /usr/app/
COPY . .

RUN apk add build-base

RUN CGO_ENABLED=1 go build -ldflags '-s -w -extldflags "-static"' -o /usr/app/appbin main.go

FROM alpine:3.14

RUN apk --update add ca-certificates && \
    rm -rf /var/cache/apk/*

ENV IMAGES_DIR /usr/app/assets/images

RUN adduser -D appuser
USER appuser

COPY --from=builder /usr/app/appbin /home/appuser/app
WORKDIR /home/appuser/

ENV DB_URL /home/appuser/the_database.db
RUN touch ${DB_URL}

COPY assets/images ${IMAGES_DIR}

CMD [ "./app" ]
