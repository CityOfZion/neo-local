<p align="center">
  <img
    src="http://res.cloudinary.com/vidsy/image/upload/v1503160820/CoZ_Icon_DARKBLUE_200x178px_oq0gxm.png"
    width="125px"
    alt="City of Zion logo">
</p>

<h1 align="center">neo-local</h1>

<p align="center">
  Quickly setup a local environment for NEO smart contract development.
</p>

<p align="center">
  <a href="https://travis-ci.org/CityOfZion/neo-local">
    <img src="https://img.shields.io/travis/CityOfZion/neo-local/master.svg">
  </a>
</p>

## Contents

1. [What?](#what)
1. [Install](#install)
1. [Usage](#usage)
1. [Block Explorer](#block-explorer)
1. [Services](#services)
1. [Troubleshooting](#troubleshooting)
1. [Credit](#credit)
1. [License](#license)

## What?

Developing smart contracts for the NEO blockchain requires a local 
[private network](https://hub.docker.com/r/cityofzion/neo-privatenet/) to be running. This project 
sets this up for you, along with a number of other utility [services](#services) to help with development.

![image](https://user-images.githubusercontent.com/2796074/36632958-9247f8ba-198d-11e8-8055-f096141902d9.png)

## Install

First start by cloning this repo to your machine:

```
git clone https://github.com/CityOfZion/neo-local.git
```

Then change to the newly cloned directory on your machine:

```
cd ./neo-local
```

Before being able to use **neo-local**, you will need to install **Docker** and **Docker Compose**.

### Mac

Use the [Docker for Mac](https://docs.docker.com/docker-for-mac/install/) installer to install both 
dependencies at the same time.

### Linux

Each dependency has to be installed separately on Linux:

1. [Docker (Community Edition)](https://store.docker.com/search?offering=community&operating_system=linux&q=&type=edition)
1. [Docker Compose](https://docs.docker.com/compose/install/#install-compose)

### Windows

Use the [Docker for Windows](https://docs.docker.com/docker-for-windows/install/) installer to install
both dependencies at the same time.

## Usage

### Mac 

Make use of the **Makefile**:

```
make start
```
```
make stop
```

### Linux

Use same commands as [Mac](#mac) users (see above).

### Windows

Windows users must run the **Docker commands manually**:

```
docker-compose up -d --build --remove-orphans
```
```
docker exec -it neo-python np-prompt -p -v
```

## Block Explorer

View what is happening on your local blockchain: [http://localhost:4000](http://localhost:4000)

## Services

The [Docker Compose](https://docs.docker.com/compose/) stack is made up of the following
services:

- [neo-privatenet](https://hub.docker.com/r/cityofzion/neo-privatenet/) (consensus nodes)
- [neo-python](https://github.com/CityOfZion/neo-python) (development CLI)
- [neo-scan](https://github.com/CityOfZion/neo-scan) (block explorer)
- [postgres](https://hub.docker.com/_/postgres/) (database for neo-scan)
- (coming sooon) [neo-faucet](https://github.com/CityOfZion/neo-faucet)

## Troubleshooting

If you have an issue then please contact any of the 
[contributors](https://github.com/CityOfZion/neo-local/graphs/contributors) on the 
[NEO Discord](https://discord.cityofzion.io), or create an [issue](https://github.com/CityOfZion/neo-local/issues/new).

The **Makefile** is designed to simplify the setup process, however in doing so it can 
obscure debugging. Thus it is recommended to run the Docker commands manually if you encounter 
an error (as outlined in [usage](#usage)).

## Credit

[@slipo](https://github.com/slipo) used Docker Compose to create a simpler
local development environment (see [repo](https://github.com/slipo/neo-scan-docker)). His work was built 
upon by the [NeoAuth](https://github.com/neoauth) team, who simplified the idea
further with the use of a wrapper and renamed the project.

It has now moved to be a part of CoZ and is actively maintained by the team - see 
[contributors](https://github.com/CityOfZion/neo-local/graphs/contributors).

**Note** - this project sits on the shoulders of a number of great CoZ projects, and wouldn't be 
possible without their hard work (see [services](#services)).

## License

[MIT](https://github.com/CityOfZion/neo-local/blob/master/LICENSE)
