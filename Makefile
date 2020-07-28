IMAGE_NAME="kube-kpifetch"
VERSION="0.1"
IMAGE="$(IMAGE_NAME):$(VERSION)"

clean:
	docker rmi -f $(shell docker images -q $(IMAGE_NAME))

build:
	docker build -t $(IMAGE) .