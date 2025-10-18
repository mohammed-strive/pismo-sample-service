FROM golang:1.25 as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o pismo-service ./main.go

FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=build /app/pismo-service .
EXPOSE 8082
CMD [ "./pismo-service" ]