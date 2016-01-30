TAG=lifthrasiir/fw.mearie.org
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
	docker run -d -p ${PORT}:80 -v /etc/ssl/certs:/etc/ssl/certs ${TAG}

.PHONY: kill
kill:
	ID=$$(docker ps -f ancestor=${TAG} -q); [ -z "$$ID" ] || docker stop $$ID

