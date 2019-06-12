# Welcome to Processbeat 1.1.0

Processbeat sends log files to Logstash or directly to Elasticsearch.

## Getting Started

To get started with Processbeat, you need to set up Elasticsearch on
your localhost first. After that, start Processbeat with:

     ./processbeat -c processbeat.yml -e

This will start Processbeat and send the data to your Elasticsearch
instance. To load the dashboards for Processbeat into Kibana, run:

    ./processbeat setup -e

For further steps visit the
[Getting started](https://www.elastic.co/guide/en/beats/libbeat/master/community-beats.html) guide.


