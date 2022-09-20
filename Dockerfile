FROM golang AS builder

ADD . /go/api
WORKDIR /go/api

RUN rm -rf deploy
RUN mkdir deploy

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o goapp ./adapter/main.go
RUN mv goapp ./deploy/goapp

FROM scratch AS production

ARG DATABASE_URL

ENV DATABASE_URL=${DATABASE_URL}
ENV PORT=3000
ENV APP_ENV="prod"

COPY --from=builder /go/api/deploy /api/

WORKDIR /api
CMD ["./goapp"]
