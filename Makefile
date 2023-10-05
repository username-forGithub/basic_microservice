app_run:
	docker compose -f docker-compose.yaml down -v
	docker compose -f docker-compose.yaml up -d uzum_delivery
	docker compose -f docker-compose.yaml up -d zookeeper
	docker compose -f docker-compose.yaml up -d kafka
	docker compose -f docker-compose.yaml up -d db
	timeout 15
	docker exec -it uzumdeliv-kafka-1 kafka-topics --create --bootstrap-server uzumdeliv-kafka-1:29092 --topic mytopic --partitions 3 --replication-factor 1


	#docker exec -it uzumdeliv-kafka-1 kafka-topics --list --bootstrap-server localhost:29092