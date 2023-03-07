# GoLang Clean Architecture

This is a project to sample Clean Architecture and Clean Clode implemented in GoLang.
Check `development` branch.

Follow the bellow steps to run the software:

1. Create a volume with the command `docker volume create cleanarchdb`
2. Run `docker-compose up -d`
3. Run the makefile command `make run-server`

Right now there are the following endpoints available:
## `GET /person/{personId}`
## `POST /person`
Body:
```
{
    "fullname": "Jane Doee",
    "age": 36,
    "email": "johndoe36@generic.com"
}
```
