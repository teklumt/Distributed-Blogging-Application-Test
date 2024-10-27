buildImage:
	docker build -t teklumt/api:auth_service0.0 .

runContainer:
	docker run --name teklumt/api:auth_service0.0 -p 8080:8080 -d auth_service

stopContainer:
	docker stop teklumt/api:auth_service0.0

removeContainer:
	docker rm teklumt/api:auth_service0.0

removeImage:
	docker rmi teklumt/api:auth_service0.0
	
logs:
	docker logs -f teklumt/api:auth_service0.0
Publish:
	docker push teklumt/api:auth_service0.0

PHONY: buildImage runContainer stopContainer removeContainer removeImage logs


