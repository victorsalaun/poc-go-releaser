FROM scratch
COPY poc-go-releaser /usr/bin/poc-go-releaser
ENTRYPOINT ["/usr/bin/poc-go-releaser"]
