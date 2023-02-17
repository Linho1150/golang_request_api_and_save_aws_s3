FROM golang:1.20-alpine3.17

ARG ACCESSKEYID_ARG
ARG ACCESSKEYSECRET_ARG
ARG API_KEY_ARG

ENV GO111MODULE=on
ENV ACCESSKEYID=$ACCESSKEYID_ARG
ENV ACCESSKEYSECRET=$ACCESSKEYSECRET_ARG
ENV API_KEY=API_KEY_ARG

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o main .
RUN go build -o repository ./repository
RUN go build -o scrapper ./scrapper

WORKDIR /dist
RUN cp /build/main .
RUN cp /build/repository .
RUN cp /build/scrapper .

FROM scratch
COPY --from=builder /dist/main .
COPY --from=builder /dist/repository .
COPY --from=builder /dist/scrapper .

ENTRYPOINT ["/main"]