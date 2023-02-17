FROM golang:1.20-alpine3.17

ARG ACCESSKEYID_ARG
ARG ACCESSKEYSECRET_ARG
ARG API_KEY_ARG

ENV ACCESSKEYID=$ACCESSKEYID_ARG
ENV ACCESSKEYSECRET=$ACCESSKEYSECRET_ARG
ENV API_KEY=API_KEY_ARG

WORKDIR /build
COPY go.mod go.sum main.go ./
RUN go mod download
RUN go build -o main .
WORKDIR /dist
RUN cp /build/main .
FROM scratch
COPY --from=builder /dist/main .
ENTRYPOINT ["/main"]