FROM library/golang AS build

MAINTAINER tailinzhang1993@gmail.com

ENV APP_DIR /go/src/locust-grpc
RUN mkdir -p $APP_DIR
WORKDIR $APP_DIR
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o locust-grpc .

# Create a minimized Docker mirror
FROM scratch AS prod

COPY --from=build /go/src/locust-grpc/locust-grpc /locust-grpc
