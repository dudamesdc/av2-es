FROM golang:latest AS builder
WORKDIR /opt/api-agendamento
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /tmp/api-agendamento github.com/dudamesdc/av2-es . 

FROM alpine:latest
WORKDIR /opt/api-agendamento
COPY --from=builder /tmp/api-agendamento .
RUN chmod +x ./api-agendamento
ENV TZ=America/Recife
RUN apk add --no-cache tzdata
COPY .env .
CMD ["./api-agendamento"]
EXPOSE 8080
