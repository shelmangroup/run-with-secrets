FROM gcr.io/distroless/base-debian10
COPY run-with-secrets /
ENTRYPOINT ["/run-with-secrets"]
