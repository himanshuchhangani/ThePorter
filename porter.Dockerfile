FROM alpine:3.14
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

RUN wget -O /usr/local/bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64
RUN chmod +x /usr/local/bin/dumb-init

# Copy library binary at the /usr/bin dir
COPY --chown=root:root cmd/main /usr/bin

# Runs "/usr/bin/dumb-init -- /my/script --with --args"
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["main"]