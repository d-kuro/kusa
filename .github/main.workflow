workflow "lint" {
  on = "push"
  resolves = ["GolangCI-Lint"]
}

action "GolangCI-Lint" {
  uses = "docker://golang:1.11.4-alpine3.8"
  runs = "sh -l -c"
  args = ["wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.12.5; which golangci-lint; which go; ./bin/golangci-lint --version;"]
}
