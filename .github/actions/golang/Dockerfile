FROM golang:1.11.4-alpine3.8

ENV GOLANGCI_LINT_VERSION="v1.12.5"

RUN wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s ${GOLANGCI_LINT_VERSION} \
  && golangci-lint --version

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
