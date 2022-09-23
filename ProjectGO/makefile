build:
	docker build --rm -t cat-api-web .
	docker image prune --filter label=stage=builder -f
	docker rmi golang:1.17-alpine3.16 alpine:3.16
run:
	docker run --rm --name cat-api-web -p 8080:8080 cat-api-web