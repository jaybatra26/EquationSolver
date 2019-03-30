dpl ?= deploy.env
include $(dpl)


build-push: build-go docker-login build-docker tag-docker docker-push
run-local: build-go build-docker run

build-go: ##build main.go
	go build main.go

docker-login: ##Login into docker hub
	docker login --username=$(DOCKER_USER)

build-docker: ## Build the container
	docker build -t $(APP_NAME) .

build-nc: ## Build the container without caching
	docker build --no-cache -t $(APP_NAME) .

tag-docker: ##Tag docker image
	docker tag $(APP_NAME) $(APP_NAME):latest

docker-push: ## push docker image
	docker push $(APP_NAME)

run: ## Run container in detatched mode on port 8080 locally
	docker run -d -p=$(PORT):$(PORT) $(APP_NAME)

run-test: ## Run unit tests
	go test
