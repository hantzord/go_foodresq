FROM golang:1.22.2 AS build-stage

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /goapp

FROM gcr.io/distroless/base-debian11 AS build-release-stage 

WORKDIR /

COPY --from=build-stage /goapp /goapp

EXPOSE 8080

USER nonroot:nonroot

CMD [ "./goapp" ]