mysql_auth:
	kubectl exec -it mysql-d49855f65-jx7mh -- mysql -u root -p

mysql_local:
	kubectl port-forward pod/mysql-d49855f65-jx7mh 3307:3306
Rabbit_local:
	kubectl port-forward service/rabbitmq 5672:5672 15672:15672
PHONY: mysql_auth