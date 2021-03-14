docker run --rm --name=logstash \
--privileged=true \
-v /home/haohao/go/src/github.com/piggona/fundingsService/logstash/pipeline:/usr/share/logstash/pipeline/ \
-v /home/haohao/go/src/github.com/piggona/fundingsService/logstash/logstash.yml:/usr/share/logstash/config/logstash.yml \
--net mini_cluster_elastic \
-d \
logstash:7.5.0
