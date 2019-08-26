COMPOSE                := docker-compose
COMPOSE_FILE           ?= config/backend.yaml

GOBIN                  := go

KAFKACAT_DOCKER_IMAGE  ?= confluentinc/cp-kafkacat
KAFKA_DOCKER_IMAGE     ?= confluentinc/cp-kafka
TAG                    ?= 5.2.3

PARTITION ?= 0
TOPIC     ?= benchmark

default: build

clean:
	@rm -rf bin/benchmark

build: clean
	@$(GOBIN) build -o ./bin/benchmark cmd/benchmark/main.go

deps: check-env
	@MY_IP=$(MY_IP) $(COMPOSE) -f $(COMPOSE_FILE) up -d

create-topic:
	$(COMPOSE) -f config/backend.yaml run --rm kafka-1 kafka-topics --create --topic $(TOPIC) --partitions 4 --replication-factor 2 --if-not-exists --zookeeper zookeeper-2:32181

listen:
	$(COMPOSE) -f config/kafkacat.yaml run --rm kafkacat kafkacat -C -b kafka-1:19092,kafka-2:29092,kafka-3:39092 -t $(TOPIC) -p $(PARTITION)

publish:
	$(COMPOSE) -f config/kafkacat.yaml run --rm kafkacat bash -c "echo 'publish to partition $(PARTITION)' | kafkacat -P -b kafka-1:19092,kafka-2:29092,kafka-3:39092 -t $(TOPIC) -p $(PARTITION)"

check-env:
ifndef MY_IP
	$(error MY_IP is not defined)
endif