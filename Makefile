TAG=lifthrasiir/fw.mearie.org
NAME=fw.mearie.org
PORT=48001

.PHONY: all
all: main

main: main.go
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o main .

.PHONY: docker
docker:
	docker build --rm --tag=${TAG} .

.PHONY: run
run: main docker
	-docker rm ${NAME}-old
	-docker rename ${NAME} ${NAME}-old && docker stop ${NAME}-old
	docker run -d -p 127.0.0.1:${PORT}:80 -v /etc/ssl/certs:/etc/ssl/certs ${TAG}
	-docker rm ${NAME}-old

