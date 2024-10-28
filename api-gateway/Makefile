build:
    docker build -t teklumt/api:api-gatewayV0.0 .

push:
    docker push teklumt/api:api-gatewayV0.0

deploy:
    kubectl apply -f kubernetes/deployment.yaml
    kubectl apply -f kubernetes/service.yaml

PHONY: build push deploy