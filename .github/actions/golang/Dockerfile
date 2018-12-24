FROM golang:1.11.4

LABEL "name"="GolangCI-Lint action"
LABEL "maintainer"="Dai Kurosawa <kurosawa7620@gmail.com>"
LABEL "version"="0.0.1"

LABEL "com.github.actions.icon"="code"
LABEL "com.github.actions.color"="green-dark"
LABEL "com.github.actions.name"="golangci-lint"
LABEL "com.github.actions.description"="This is an action to run GolangCI-Lint"

ENV GOLANGCI_LINT_VERSION="v1.12.3"

RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh \
  | sh -s -- -b $(go env GOPATH)/bin ${GOLANGCI_LINT_VERSION}

ADD entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
