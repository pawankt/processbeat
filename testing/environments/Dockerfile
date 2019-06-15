# Basic debian file with curl, wget and nano installed to fetch files
# an update config files
FROM debian:latest

RUN apt-get update && \
    apt-get install -y curl nano wget zip && \
    apt-get clean
	apt-get virtualenv
	apt-get python-pip
	apt-get netcat
	apt-get libpcap-dev
	
RUN pip install --upgrade pip
RUN pip install --upgrade setuptools
RUN pip install --upgrade docker-compose==1.23.2
