# NEO private network - Dockerfile
FROM microsoft/dotnet:2.1.4-runtime-bionic

LABEL maintainer="City of Zion"
LABEL authors="original authors metachris, ashant, hal0x2328, phetter"

# arguments to choose version of neo-cli to install (defaults to 2.8.0)
ARG NEO_CLI_VERSION="2.8.0"

# Frontend non-interactive
ENV DEBIAN_FRONTEND noninteractive

# Disable dotnet usage information collection
# https://docs.microsoft.com/en-us/dotnet/core/tools/telemetry#behavior
ENV DOTNET_CLI_TELEMETRY_OPTOUT 1

# Install system dependencies. always should be done in one line
# https://docs.docker.com/engine/userguide/eng-image/dockerfile_best-practices/#run
RUN apt-get update && apt-get install -y \
    unzip \
    screen \
    expect \
    libleveldb-dev \
    wget \
    curl \
    git-core \
    python3.6 \
    python3.6-dev \
    python3.6-venv \
    python3-pip \
    libssl-dev \
    vim \
    man \
    libunwind8

# APT cleanup to reduce image size
RUN rm -rf /var/lib/apt/lists/*

# Download, add and decomnpres the neo-cli package. At the end, delete the zip file.
# $NEO_CLI_VERSION is a build argument
RUN wget -O /opt/neo-cli.zip "https://github.com/neo-project/neo-cli/releases/download/v$NEO_CLI_VERSION/neo-cli-linux-x64.zip" && \
    unzip -q -d /opt/node /opt/neo-cli.zip && \
    rm /opt/neo-cli.zip

# Download, add and decomnpres the simplePoplicy plugin package. At the end, delete the zip file.
RUN wget -O /opt/SimplePolicy.zip "https://github.com/neo-project/neo-plugins/releases/download/v2.9.0/SimplePolicy.zip" && \
    unzip -q -d /opt/node/neo-cli /opt/SimplePolicy.zip && \
    rm /opt/SimplePolicy.zip

# neo-python setup: clonse and install dependencies
RUN git clone https://github.com/CityOfZion/neo-python.git /neo-python
WORKDIR /neo-python
# RUN git checkout development
RUN pip3 install -e .

# Add config files
ADD ./configs/v$NEO_CLI_VERSION/* /opt/node/neo-cli/
ADD ./wallets/* /opt/node/neo-cli/

# Add scripts
RUN wget https://s3.amazonaws.com/neo-experiments/neo-privnet.wallet
ADD ./scripts/run.sh /opt/
ADD ./scripts/start_consensus_node.sh /opt/
ADD ./scripts/claim_neo_and_gas_fixedwallet.py /neo-python/
ADD ./scripts/claim_gas_fixedwallet.py /neo-python/
ADD ./wallets/neo-privnet.python-wallet /tmp/wallet

# A welcome message for bash users
RUN echo "printf \"\n* Consensus nodes are running in screen sessions, check 'screen -ls'\"" >> /root/.bashrc
RUN echo "printf \"\n* Please report issues to https://github.com/CityOfZion/neo-local\n\n\"" >> /root/.bashrc

# On docker run, start the consensus nodes
CMD ["/bin/bash", "/opt/run.sh"]