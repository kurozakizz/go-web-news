dev: 
	go run cmd/server.go

format:
	gofmt -w .

commit:
	git commit -am "$m"

push:
	git push origin master

all:
	$(eval GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD))
	echo Git branch is $(GIT_BRANCH)
