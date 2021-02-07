PATH := $(PATH):/usr/local/go/bin
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GODEP=$(GOCMD) mod
BINARY_NAME=rds-events-exporter
BINARY_DIR=./bin
DOCKER=docker
IMAGE_NAME=aws-exporter


all: test build
build:
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME)

test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_DIR)/$(BINARY_NAME)
run: build
	$(BINARY_DIR)/$(BINARY_NAME)
deps:
	$(GODEP) download

docker-build:
	$(DOCKER) build . -t $(IMAGE_NAME):latest

tag-version:
	@echo 'create tag $(VERSION)'
	$(DOCKER) tag $(IMAGE_NAME) $(DOCKER_REPO)/$(IMAGE_NAME):$(VERSION)

tag-latest:
	@echo 'create tag latest'
	$(DOCKER) tag $(IMAGE_NAME) $(DOCKER_REPO)/$(IMAGE_NAME):latest

tag: tag-latest tag-version

publish-version:
	@echo 'publish $(VERSION) to $(DOCKER_REPO)'
	docker push $(DOCKER_REPO)/$(IMAGE_NAME):$(VERSION)

publish-latest:
	@echo 'publish latest to $(DOCKER_REPO)'
	docker push $(DOCKER_REPO)/$(IMAGE_NAME):latest

repo-login:
	@echo 'Logging in to $(DOCKER_REPO)'

publish: tag repo-login publish-latest publish-version