docker run -d --rm --name kafka --publish 9092:9092 \
--link zookeeper \
--env KAFKA_ZOOKEEPER_CONNECT=172.30.39.100:2181 \
--env KAFKA_ADVERTISED_HOST_NAME=172.30.39.100 \
--env KAFKA_ADVERTISED_PORT=9092  \
wurstmeister/kafka