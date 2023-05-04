FROM mongo:6.0.5-jammy

ENV TZ="Asia/Tokyo"
ENV MONGO_INITDB_ROOT_USERNAME root
ENV MONGO_INITDB_ROOT_PASSWORD brHZ-!_rHAZF4xR2-EsRKx9e
RUN echo $TZ > /etc/timezone

COPY init.js /docker-entrypoint-initdb.d/