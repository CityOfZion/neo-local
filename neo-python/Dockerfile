FROM python:3.5
LABEL maintainer="NeoAuth <hi@neoauth.org>"

RUN apt-get update
RUN apt-get install libleveldb-dev libssl-dev g++ unzip -y

RUN wget https://github.com/CityOfZion/neo-python/archive/development.zip
RUN unzip development.zip

WORKDIR neo-python-development
RUN sed -i -e 's/127.0.0.1/neo-nodes/g' protocol.privnet.json

RUN wget https://s3.amazonaws.com/neo-experiments/neo-privnet.wallet
RUN pip install -e .

CMD bash