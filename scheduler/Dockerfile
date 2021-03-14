FROM debian:stretch
COPY ./scheduler /
ENV BROKERS kafka:9092
ENV DB_PASSWD vanlink
ENV DB_HOST mysql:3306
RUN mkdir /app \
    && mv /scheduler /app/
ENTRYPOINT ["./scheduler"]
WORKDIR /app