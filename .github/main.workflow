workflow "lint" {
  on = "push"
  resolves = ["GolangCI-Lint"]
}

action "GolangCI-Lint" {
  uses = "./.github/actions/golang"
}
