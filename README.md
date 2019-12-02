# ethclient
Basic ethereum block explorer


## Required Environment Vars
```
export HTTP_PORT=8080
export ENVIRONMENT=mainnet
```

Environment can be any environment infura supports
https://infura.io/docs/gettingStarted/chooseaNetwork

mainnet
rinkeby
ropsten


## Running in docker
```
docker build -t chadgrant/ethclient .
docker run -itp 8080:8080 --rm --env-file docker.env chadgrant/ethclient
```

## Building / Running locally
```
export HTTP_PORT=8080
export ENVIRONMENT=mainnet

go get ./...
go build
./ethclient
```

Once server is running, you can browse to http://localhost:8080/
