buildImage:
	docker build -t teklumt/api:userservice0.0 .

runContainer:
	docker run --name teklumt/api:userservice0.0 -p 8080:8080 -d userservice

stopContainer:
	docker stop teklumt/api:userservice0.0

removeContainer:
	docker rm teklumt/api:userservice0.0

removeImage:
	docker rmi teklumt/api:userservice0.0
	
logs:
	docker logs -f teklumt/api:userservice0.0
Publish:
	docker push teklumt/api:userservice0.0

PHONY: buildImage runContainer stopContainer removeContainer removeImage logs


