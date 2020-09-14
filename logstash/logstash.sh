docker run --rm --name=logstash \
--privileged=true \
-v /home/haohao/fundingsService/logstash/pipeline:/usr/share/logstash/pipeline/ \
-v /home/haohao/fundingsService/logstash/logstash.yml:/usr/share/logstash/config/logstash.yml \
--net cluster_elastic \
-d \
docker.elastic.co/logstash/logstash:7.5.0
