docker run --name es01 \
-v esdata01:/usr/share/elasticsearch/data \
-v esdata01log:/usr/share/elasticsearch/logs \
--net es_vanlink \
--publish-all \
-e ES_JAVA_OPTS="-Xms512m -Xmx512m" \
-e TAKE_FILE_OWNERSHIP=true \
--ulimit memlock=-1:-1 \
-d \
docker.elastic.co/elasticsearch/elasticsearch:7.5.0



