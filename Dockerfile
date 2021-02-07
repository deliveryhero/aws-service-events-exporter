FROM golang:1.15.6-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
RUN apk --no-cache add ca-certificates
COPY ./go.mod .
RUN go mod download -x
COPY . .
RUN go build -o /app .

FROM scratch AS final
COPY --from=build /etc/ssl/certs /etc/ssl/certs
COPY --from=build /app /app
ENTRYPOINT ["/app"]