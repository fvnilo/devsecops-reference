FROM golang:1.24-alpine AS build

WORKDIR /app

RUN mkdir -p libs apps

COPY ./libs/reverse_go ./libs/reverse_go
COPY ./apps/service-b ./apps/service-b

RUN go work init ./apps/service-b ./libs/reverse_go
RUN go build -o server ./apps/service-b

FROM alpine:3.22

WORKDIR /app
COPY --from=build /app/server .

EXPOSE 8081

CMD ["./server"]
