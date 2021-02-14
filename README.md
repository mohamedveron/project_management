# project_management
This app is developed using go-chi https://github.com/go-chi/chi router and oapi-codegen https://github.com/deepmap/oapi-codegen to generate client for apis

# Use api.yml file in the swagger ui to see the description of the rest apis


## Covered Scenarios by the Unit tests :

1- Create New Project

2- Update New Project. 

3- Assign Project Owner with role manage.

4- Trying to Assign Project Owner with not manage role.

5- Assign Project Participants 

6- Trying Assign Project Participants with not the same role as the owner

## Setup

Must have golang installed version >= 12.0.0

make file consists of 4 steps: generate, test, build, run
you can run all of them 

```bash
make all
```

# Run the unit tests:
```bash
make test
```
If you have issue with code generation in generate step you can run go test -v ./tests

# Start the http server:

```bash
make run
```

If you have issue with code generation in generate step you can copy the api/api.gen.go file in repo and run go run main.go to start the server

### Notes:

1- I didn't use any database just in memory hash maps for both projects and employees

2- There is an list employees api

