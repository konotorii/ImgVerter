FROM cgr.dev/chainguard/static@sha256:288b818c1b3dd89776d176f07f5f671b118fe836c4d80ec2cc3299b596fe71b7
ADD https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.5.0-linux-x86-64.tar.gz /bin/webp/libwebp.tar.gz
ADD --chmod=0755 /bin/webp/libwebp.tar.gz /bin/webp
COPY imgverter \
	/usr/bin/imgverter
ENV SKIP_DOWNLOAD=true
ENV VENDOR_PATH=/bin/webp
ENTRYPOINT ["/usr/bin/imgverter"]