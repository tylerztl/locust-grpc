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
