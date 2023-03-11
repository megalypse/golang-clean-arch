# GoLang Clean Architecture

This is a project to sample Clean Architecture and Clean Code implemented in GoLang.
Check `development` branch.

When first clonning the app, you will get an error at `internal/main/factory/router_factory.go`,
but it's normal, just follow the bellow steps:

---
These steps requires `github.com/swaggo/swag/cmd/swag` to be installed in your `$GOBIN`, which should be already added
to your `$PATH`.

You can check if your `$GOBIN` directory is in your `$PATH` by running the following command:
`echo $PATH | grep $(go env GOPATH)/bin`

If this command does not output anything, it means that your `$GOBIN` directory is not in your `$PATH`.
To fix this, you can add your `$GOBIN` directory to your `$PATH` by adding the following line to your shell configuration
file (e.g. `~/.bashrc` or `~/.zshrc`):
`export PATH=$PATH:$(go env GOPATH)/bin`

---


1. Run `make run-compose`

Check `http://localhost:3001/swagger/index.html` for Swagger Documentation.

---

# Makefile commands
* `clean`: Clean server's build
* `generate-docs`: Generates new swagger docs
* `run-compose`: Run compose without cleanup
* `run-compose-clean`: Clean up server build and run compose with a new one
* `run-compose-clean-all`: Clean up server and database and do a fresh run of compose
