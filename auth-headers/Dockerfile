FROM golang:1.12.5-stretch@sha256:d17a1d8f0c20d108d1177d560f4afb9de10104c46df756d885cfa4282bbaac65

WORKDIR /src

COPY go.* ./
RUN go mod download

COPY . .

RUN go test

RUN sed -i 's/zap.NewDevelopment()/zap.NewProduction()/' main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags '-w -extldflags "-static"'

FROM scratch

COPY --from=0 /src/auth-headers /usr/local/bin/auth-headers

ENV LISTEN=":3001"

ENTRYPOINT ["auth-headers"]
