buildImage:
	docker build -t teklumt/api:api-gatewayV0.0 .

runContainer:
	docker run -d -p 8080:8080 --name teklumt/api:api-gatewayV0.0 api-gateway

stopContainer:
	docker stop teklumt/api:api-gatewayV0.0

removeContainer:
	docker rm teklumt/api:api-gatewayV0.0

pushImage:
	docker push teklumt/api:api-gatewayV0.0

PHONY: buildImage runContainer stopContainer removeContainer