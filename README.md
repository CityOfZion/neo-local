<p align="center">
  <img
    src="http://res.cloudinary.com/vidsy/image/upload/v1503160820/CoZ_Icon_DARKBLUE_200x178px_oq0gxm.png"
    width="125px"
    alt="City of Zion logo">
</p>

<h1 align="center">neo-local</h1>

<p align="center">
  Start a local NEO blockchain with a single command: <code>make start</code>
</p>

<p align="center">
  <a href="https://travis-ci.org/CityOfZion/neo-local">
    <img src="https://img.shields.io/travis/CityOfZion/neo-local/master.svg">
  </a>
</p>

## What?

Developing smart contracts for the NEO blockchain requires a local private network to be running. This project sets this up for you, along with a number of other utility services to help with development.

![image](https://user-images.githubusercontent.com/2796074/36632958-9247f8ba-198d-11e8-8055-f096141902d9.png)

## Prerequisites

Install [Docker CE](https://www.docker.com/community-edition) and [Docker Compose](https://docs.docker.com/compose/).

## Usage

You can either start the full setup with Docker commands or the Makefile:

**Docker commands**

    # Start the containers:
    docker-compose up
   
    # Run neo-python:
    docker exec -it neo-python np-prompt -p -v

**Makefile**

    # Start and attach to neo-python:
    make start
    
    # Stop
    make stop

## Block Explorer

View what is happening on your local blockchain: [http://localhost:4000](http://localhost:4000)

## Docker Compose Services

The [Docker Compose](https://docs.docker.com/compose/) stack is made up of the following
services:

- [neo-privatenet](https://hub.docker.com/r/cityofzion/neo-privatenet/) (consensus nodes)
- [neo-python](https://github.com/CityOfZion/neo-python) (development CLI)
- [neo-scan](https://github.com/CityOfZion/neo-scan) (block explorer)
- [postgres](https://hub.docker.com/_/postgres/) (database for neo-scan)
- (coming sooon) [neo-faucet](https://github.com/CityOfZion/neo-faucet)

## Credit

This repo was originally created by [@slipo](https://github.com/slipo) and then built into **neo-local** by the [NeoAuth](https://github.com/neoauth) team.

It is now moved to be a part of CoZ and is actively maintained by the team - see [contributors](https://github.com/CityOfZion/neo-local/graphs/contributors).

## License

[MIT](https://github.com/CityOfZion/neo-local/blob/master/LICENSE)

## Troubleshooting

If you have an issue then please contact any of the [contributors](https://github.com/CityOfZion/neo-local/graphs/contributors) on the [NEO Discord](https://discord.cityofzion.io).
