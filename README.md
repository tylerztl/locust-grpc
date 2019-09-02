# locust-grpc

### Prerequisites
- Go 1.10+ installation or later
- GOPATH environment variable is set correctly
- Python 2.7.10+ installation or later

## Getting started

### install locust
```
pip install locust
```

### run master
```
locust -f dummy.py --master --master-bind-host=127.0.0.1 --master-bind-port=5557
```

#### run slave
```
go run boomer.go --master-host=127.0.0.1 --master-port=5557  
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
