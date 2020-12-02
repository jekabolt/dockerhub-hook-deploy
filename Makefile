IMAGE_NAME=hook-deploy

build:
	go build -o ./bin/$(IMAGE_NAME) ./

run: build
	./bin/$(IMAGE_NAME)

dist:
	env GOOS=linux GOARCH=amd64 go build -o ./bin/$(IMAGE_NAME) ./	
	
scp: dist
	source .env && scp ./bin/$(IMAGE_NAME) ${REMOTE_SSH}:

