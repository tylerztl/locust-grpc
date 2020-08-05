# locust-grpc

### Prerequisites
- Go 1.10+ installation or later
- GOPATH environment variable is set correctly
- Python 3.8 installation

## Getting started

### install [locust](https://docs.locust.io/en/stable/index.html)
```
pip install locust
```

### run master
```
 locust --master -f dummy.py
```

#### run slave
```
go run boomer.go --max-rps 1000
```

#### web test
visit：http://127.0.0.1:8089/

## Docker started

### build locust slave image
```
docker build -t hyperledger/locust-grpc .
```

### run master
```
nohup locust -f dummy.py --master --master-bind-host=127.0.0.1 --master-bind-port=5557 > locust.log 2>&1 &
```
#### run slave
```
docker-compose up -d
```
