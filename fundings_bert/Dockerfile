FROM keywords_service_env:latest
WORKDIR /Project/service

COPY ./bert/ .
COPY ./data/ .

ENV BROKERS kafka:9092

RUN pip install kafka-python PyMySQL

VOLUME /Project/service/get_keywords/tmp

ENTRYPOINT ["python","./kafka_service.py"]