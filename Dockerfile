FROM debian:bookworm

RUN apt-get update
RUN apt-get install -y webp

ENV SKIP_DOWNLOAD=true
ENV VENDOR_PATH=/usr/bin

ENTRYPOINT ["/usr/bin/imgverter"]