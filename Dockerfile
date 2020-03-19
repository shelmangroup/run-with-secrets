FROM gcr.io/distroless/base-debian10:debug
COPY run-with-secrets /
ENTRYPOINT ["/run-with-secrets"]
