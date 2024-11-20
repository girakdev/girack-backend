# GIRACK BACKEND


### start server

```sh
go install github.com/swaggo/swag/cmd/swag@latest
go install go install github.com/golang/mock/mockgen@v1.6.0
```
```sh
make dup
```
```sh
make serve
```

### Show swagger
localhost:8080/swagger/index.html
### Test

```sh
./scripts/run-in-docker.sh go test ./... -race
```