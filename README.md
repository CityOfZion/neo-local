<p align="center">
  <img
    src="http://res.cloudinary.com/vidsy/image/upload/v1503160820/CoZ_Icon_DARKBLUE_200x178px_oq0gxm.png"
    width="125px"
    alt="City of Zion logo">
</p>

<h1 align="center">neo-local</h1>

<p align="center">
    <a href="https://hub.docker.com/r/cityofzion/neo-privatenet/">neo privatenet</a> + <a href="http://neoscan.io/">neoscan</a> + <a href="https://neo-python.readthedocs.io">neopython</a>
</p>
<p align="center">
    Start a local NEO blockchain with a single command: `make start`
</p>

<p align="center">
  <a href="https://travis-ci.org/CityOfZion/neo-local">
    <img src="https://img.shields.io/travis/CityOfZion/neo-local/master.svg">
  </a>
</p>

## What?

Developing smart contracts requires a local private network of the NEO blockchain to be running. This projects has everything you need to get started with `make start`:

* [neo-privatenet](https://hub.docker.com/r/cityofzion/neo-privatenet/) ([GitHub](https://github.com/CityOfZion/neo-privatenet-docker))
* [neoscan](http://neoscan.io/) ([GitHub](https://github.com/CityOfZion/neo-scan))
* [neo-python](https://neo-python.readthedocs.io) (for convenience) ([GitHub](https://github.com/CityOfZion/neo-python))
* (hopefully soon, [neo-faucet](https://github.com/CityOfZion/neo-faucet))

[Docker Compose](https://docs.docker.com/compose/) is used to package and run all of the services.

![image](https://user-images.githubusercontent.com/2796074/36632958-9247f8ba-198d-11e8-8055-f096141902d9.png)

## Usage

Install [Docker Compose](https://docs.docker.com/compose/), then run:

```
make start
```

To stop:

```
make stop
```

## Block Explorer

View what is happening on your local blockchain: [http://localhost:4000](http://localhost:4000)

## Docker Compose Swrvices

The [Docker Compose](https://docs.docker.com/compose/) stack is made up of the following
services:

- [neo-privatenet](https://hub.docker.com/r/cityofzion/neo-privatenet/) (consensus nodes)
- [neo-python](https://github.com/CityOfZion/neo-python) (development CLI)
- [neo-scan](https://github.com/CityOfZion/neo-scan) (block explorer)
- [postgres](https://hub.docker.com/_/postgres/) (database for neo-scan)

## Credit

This is a fork of [@slipo](https://github.com/slipo)'s
[neo-scan-docker](https://github.com/slipo/neo-scan-docker) repo and [NeoAuth/neo-local](https://github.com/neoauth/neo-local).

[Contributors](https://github.com/CityOfZion/neo-local/graphs/contributors)

## License

MIT
