FROM golang:1.24-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN cd cmd/server && go build -o /bin/bodego

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /bin/bodego /bin/bodego
EXPOSE 8080
CMD ["/bin/bodego"]
