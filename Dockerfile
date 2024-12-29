# FROM golang:1.22-bookworm
# ENV CGO_ENABLED=1

# RUN apt-get update \
#  && DEBIAN_FRONTEND=noninteractive \
#     apt-get install --no-install-recommends --assume-yes \
#       build-essential \
#       libwebp-dev

# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o /imgverter ./

# FROM debian:bookworm
# RUN apt-get update \
#  && DEBIAN_FRONTEND=noninteractive \
#     apt-get install --no-install-recommends --assume-yes \
#       libwebp7
# COPY --from=0 imgverter /usr/bin/imgverter
# EXPOSE 8000
# CMD ["/usr/bin/imgverter"]

FROM cgr.dev/chainguard/static@sha256:288b818c1b3dd89776d176f07f5f671b118fe836c4d80ec2cc3299b596fe71b7
COPY imgverter \
	/usr/bin/imgverter
ENTRYPOINT ["/usr/bin/imgverter"]