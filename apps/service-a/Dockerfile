FROM golang:1.24-alpine AS build

WORKDIR /app

RUN mkdir -p libs apps

COPY ./apps/service-a ./apps/service-a

RUN go work init ./apps/service-a
RUN go build -o server ./apps/service-a

FROM alpine:3.22

WORKDIR /app
COPY --from=build /app/server .

EXPOSE 8080

CMD ["./server"]
