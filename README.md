# docker-control

run or rm a docker container via go rest api.

Cross-domain issues have been already resolved in this code.

#### first u need to build go_rest:

```bash
env GOOS=linux GOARCH=amd64 go build -o go_rest go_rest.go
```

#### then:

```bash
nohup ./gorest &
```
