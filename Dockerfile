FROM golang AS builder

ADD . /go/api
WORKDIR /go/api

RUN rm -rf deploy
RUN mkdir deploy
ARG DATABASE_URL

ENV DATABASE_URL=${DATABASE_URL}
ENV PORT=3000
ENV APP_MODE="prod"

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o goapp ./adapter/main.go
RUN mv goapp ./deploy/goapp
RUN mv .env ./deploy/

FROM scratch AS production
COPY --from=builder /go/api/deploy /api/

WORKDIR /api
ENTRYPOINT ./goapp
