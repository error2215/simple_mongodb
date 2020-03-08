# All dependencies
mod:
	@echo "======================================================================"
	@echo "Run MOD"
	@ GO111MODULE=on go mod verify
	@ GO111MODULE=on go mod tidy
	@ GO111MODULE=on go mod vendor
	@ GO111MODULE=on go mod download
	@ GO111MODULE=on go mod verify
	@echo "======================================================================"

docker:
	@echo "Docker-compose up"
	@docker-compose up -d

docker-build:
	@echo "Docker-compose build"
	@docker-compose build

all_docker: mod build docker

mongo_docker:
	@sudo docker run -d -p 27017:27017 -v ~/data:/data/db mongo