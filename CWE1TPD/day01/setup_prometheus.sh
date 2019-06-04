#!/bin/bash
docker run -d -p 9090:9090 -v /tmp/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
docker run -d -p 9100:9100 --net="host" prom/node-exporter
