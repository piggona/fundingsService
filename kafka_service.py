from kafka import KafkaConsumer
from kafka import TopicPartition
from kafka import KafkaProducer
import service
import time
import json
import os

my_topic = 'import_raw'
brokers = ['172.30.39.100:9092']
if os.getenv('TOPIC') != '':
    my_topic = os.getenv('TOPIC')
if os.getenv('BROKERS') != '':
    brokers = os.getenv('BROKERS').split(';')

def start_consumer():
    consumer = KafkaConsumer( bootstrap_servers = brokers)
    producer = KafkaProducer(bootstrap_servers=brokers,key_serializer= str.encode, value_serializer= str.encode)
    consumer.assign([
        TopicPartition(topic=my_topic,partition=0)
    ])
    consumer.seek(partition=TopicPartition(topic=my_topic,partition=0),offset=345)
    for msg in consumer:
        print(msg)
        print("topic = %s" % msg.topic) # topic default is string
        print("partition = %d" % msg.offset)
        print("value = %s" % msg.value.decode()) # bytes to string
        print("timestamp = %d" % msg.timestamp)
        print("time = ", time.strftime("%Y-%m-%d %H:%M:%S", time.localtime( msg.timestamp/1000 )) )
        fund_content = json.loads(msg.value.decode())
        keywords = service.GetKeywords()
        results = keywords.operate(fund_content['description'])
        industries = []
        technology = []
        for result in results:
            if result['label'] == 'indu':
                industries.append(result['text'])
            elif result['label'] == 'tech':
                technology.append(result['text'])
        fund_content['industries'] = industries
        fund_content['technology'] = technology
        future = producer.send(my_topic ,  key= 'import_raw', value= json.dumps(fund_content), partition= 0)
        future.get(timeout=10)

if __name__ == '__main__':
    start_consumer()