# GoLang Clean Architecture

This is a project to sample Clean Architecture and Clean Code implemented in GoLang.
Check `development` branch.

When first clonning the app, you will get an error at `internal/main/factory/router_factory.go`,
but it's normal, just follow the bellow steps:

1. Create a volume with the command `docker volume create cleanarchdb`
2. Run `make run-compose`

Check `http://localhost:3001/swagger/index.html` for Swagger Documentation.

# Makefile commands
* `clean`: Clean server's build
* `generate-docs`: Generates new swagger docs
* `run-compose`: Run compose without cleanup
* `run-compose-clean`: Clean up server build and run compose with a new one
* `run-compose-clean-all`: Clean up server and database and do a fresh run of compose
