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
	docker build --target prod -t "${registry_url}/${username}/${repository_name}:${version}" -f Dockerfile .
	docker tag "${registry_url}/${username}/${repository_name}:${version}" "${registry_url}/${username}/${repository_name}:latest"
	echo "${password}" | docker login -u "${username}" --password-stdin "${registry_url}"
	docker push "${registry_url}/${username}/${repository_name}:${version}"
	docker push "${registry_url}/${username}/${repository_name}:latest"
	docker logout "${registry_url}"

test:
	echo "Building test container"
	docker build --target dev-builder --progress=plain -t "${registry_url}/${username}/${repository_name}:test" -f Dockerfile .
