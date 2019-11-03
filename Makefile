build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o repo-init ./src/

build-docker:
	docker run --rm -v $$PWD:/code -e CGO_ENABLED=0 -e GOOS=linux -w /code golang:1.13 go build -a -installsuffix cgo -o repo-init ./src/
