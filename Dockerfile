FROM centurylink/ca-certs
COPY botuser /
ENTRYPOINT ["/botuser"]
