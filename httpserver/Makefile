APP_NAME=httpserver
ADDR=80
VERSION="v1.0.0"

TAG="v1.0"
DOCKER_ACCOUNT="bearlu"

run: build
	@echo "run app..."
	@chmod +x ./$(APP_NAME)
	@sudo ./$(APP_NAME) -addr=:$(ADDR) -version=$(VERSION)

build:
	@echo "build image..."
	@docker build -t $(DOCKER_ACCOUNT)/$(APP_NAME):$(TAG) -f ./Dockerfile .

push: build
	@echo "push image $(APP_NAME)..."
	@docker tag $(DOCKER_ACCOUNT)/$(APP_NAME):$(TAG) $(DOCKER_ACCOUNT)/$(APP_NAME):$(TAG)
	@docker push $(DOCKER_ACCOUNT)/$(APP_NAME):$(TAG)

run-image-local: build
	@echo "run image local..."
	@docker run -p $(ADDR):$(ADDR) $(DOCKER_ACCOUNT)/$(APP_NAME) -addr=:$(ADDR) -version=$(VERSION)

run-image:
	@echo "run image..."
	@docker run -p $(ADDR):$(ADDR) $(DOCKER_ACCOUNT)/$(APP_NAME):$(TAG) -addr=:$(ADDR) -version=$(VERSION)


.PHONY: \
		build \
		run-image-local \
		run-image \
		push \
		run
