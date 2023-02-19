FROM public.ecr.aws/lambda/provided:al2 as build
LABEL email="linho301150@gmail.com"

ARG ACCESSKEYID_ARG
ARG ACCESSKEYSECRET_ARG
ARG API_KEY_ARG

ENV GO111MODULE on
ENV ACCESSKEYID $ACCESSKEYID_ARG
ENV ACCESSKEYSECRET $ACCESSKEYSECRET_ARG
ENV API_KEY $API_KEY_ARG

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN yum install -y golang
RUN go env -w GOPROXY=direct

ADD go.mod go.sum ./
RUN go mod download

ADD . .
RUN go build -o /linho1150

FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /linho1150 /linho1150
ENTRYPOINT ["/linho1150"]