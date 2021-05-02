FROM debian:buster-slim

RUN apt update && apt install -y ca-certificates

COPY out/* .
CMD ["./app"]
