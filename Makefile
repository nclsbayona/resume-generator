##########
# Variable version should be passed as an argument to the make command as follows: make version=1.0.0 release
# Variable password should be passed as an argument to the make command as follows: make password=admin
# Variable registry_url should be passed as an argument to the make command as follows: make registry_url=docker.io
# Variable repository_name should be passed as an argument to the make command as follows: make repository_name=repository
##########
username := $(shell echo $(repository_name) | awk -F/ '{print $1}')
release: test
	echo "New release"
	echo "Release of version: ${version}"
	echo "${password}" | docker login -u "${username}" --password-stdin "${registry_url}"
	docker buildx build --progress=plain --target prod --tag "${registry_url}/${username}/${repository_name}:${version}" --push --platform linux/amd64,linux/arm/v7,linux/arm64 .
	docker buildx build --progress=plain --target prod --tag "${registry_url}/${username}/${repository_name}:latest" --push --platform linux/amd64,linux/arm/v7,linux/arm64 .
	docker logout "${registry_url}"

test:
	echo "Building test container"
	docker buildx build --progress=plain --target dev-builder --tag test --platform linux/amd64,linux/arm/v7,linux/arm64 .
