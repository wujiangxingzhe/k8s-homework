BIN_FILE="httpserver"
GO_PROJEDT_DIR="/opt/go_workspace/k8s-homework-master/"
DOCKER_CONTEXT="/opt/docker"
DOCKER_FILE="Dockerfile"
REPO="wujiangxingzhe"
TAG="v0.1"

.PHONY: build docker_build docker_push docker_run

build: 
	sh /opt/scripts/build_httpserver.sh "${GO_PROJEDT_DIR}" "${BIN_FILE}" "${DOCKER_CONTEXT}"

docker_build: build
	docker build -f "${DOCKER_CONTEXT}"/"${DOCKER_FILE}" "${DOCKER_CONTEXT}" -t "${REPO}"/"${BIN_FILE}":"${TAG}"

docker_run: docker_build
	docker rm -f "${BIN_FILE}" && docker run --name "${BIN_FILE}" -d "${REPO}"/"${BIN_FILE}":"${TAG}"

docker_push: docker_build
	docker login && docker push "${REPO}"/"${BIN_FILE}":"${TAG}"
