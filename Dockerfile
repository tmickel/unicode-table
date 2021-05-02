FROM debian:buster-slim

RUN apt install -y ca-certificates

COPY out/* .
CMD ["./app"]
