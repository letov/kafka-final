docker_bin := $(shell command -v docker 2> /dev/null)
docker_compose_bin := $(shell command -v docker-compose 2> /dev/null)

up:
	$(docker_compose_bin) up -d

down:
	$(docker_compose_bin) down -v

restart: down up

ksqldb-cli:
	$(docker_compose_bin) exec ksqldb-cli ksql http://ksqldb-server:8088

init: create-topics create-connectors ksqldb-migrations

create-topics: delete-topics
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics --create --topic shop_products --bootstrap-server kafka1:9092 --partitions 2 --replication-factor 1
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics --create --topic analytic_products_filtered --bootstrap-server kafka1:9092 --partitions 2 --replication-factor 1
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics --create --topic analytic_products_find --bootstrap-server kafka1:9092 --partitions 2 --replication-factor 1

delete-topics:
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics --if-exists --delete --topic analytic_products_find --bootstrap-server kafka1:9092
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics --if-exists --delete --topic analytic_products_filtered --bootstrap-server kafka1:9092
	$(docker_compose_bin) exec -it kafka1 ../../usr/bin/kafka-topics --if-exists --delete --topic shop_products --bootstrap-server kafka1:9092

create-connectors: delete-connectors
	curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" http://localhost:8083/connectors/ -d @docker/postgres-connector.json

delete-connectors:
	curl -X DELETE http://localhost:8083/connectors/postgres-connector

ksqldb-migrations:
	$(docker_compose_bin) exec ksqldb-cli ksql http://ksqldb-server:8088 -f '/docker/ksql-queries.sql'

ksqldb-rollback:
	$(docker_compose_bin) exec ksqldb-cli ksql http://ksqldb-server:8088 -f '/docker/ksql-rollback.sql'