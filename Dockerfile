FROM golang:latest AS builder

WORKDIR /app
COPY . .
RUN go mod download && go build -ldflags "-s -w" -o 1kv-exporter cmd/1kv-exporter/main.go


FROM debian:latest AS stager

ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update && apt-get -y upgrade && apt-get -y install ca-certificates
RUN groupadd -g 65510 exporter && useradd -u 65510 -m -d /exporter -g exporter -r exporter


FROM scratch

COPY --from=stager /etc/passwd /etc/passwd
COPY --from=stager /etc/group /etc/group
COPY --from=stager --chown=exporter:exporter /exporter /exporter
COPY --from=stager /usr/lib /usr/lib
COPY --from=stager /lib /lib
COPY --from=stager /lib64 /lib64
COPY --from=stager /etc/ssl /etc/ssl

COPY --from=builder /app/1kv-exporter /

EXPOSE 17586
WORKDIR /exporter
USER exporter

ENTRYPOINT ["/1kv-exporter"]
