# Kubernetes
deploy:
	@kubectl apply -f kubernetes/role.yaml
	@kubectl apply -f kubernetes/server-deployment.yaml
	@kubectl apply -f kubernetes/server-service.yaml
	@kubectl apply -f kubernetes/client-deployment.yaml
	@kubectl apply -f kubernetes/publisher-deployment.yaml

deploy-kafka:
	@kubectl apply -f kubernetes/kafka

destroy:
	@kubectl delete -f kubernetes/publisher-deployment.yaml
	@kubectl delete -f kubernetes/client-deployment.yaml
	@kubectl delete -f kubernetes/server-service.yaml
	@kubectl delete -f kubernetes/server-deployment.yaml
	@kubectl delete -f kubernetes/role.yaml

destroy-kafka:
	@kubectl delete -f kubernetes/kafka
