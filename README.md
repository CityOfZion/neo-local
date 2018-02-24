<p align="center">
  <img src="https://raw.githubusercontent.com/neoauth/design-assets/master/logo/colour/neoauth_color.png" width="150px" /> 
</p>

<h1 align="center">neo-local</h1>

<p align="center">
  Start a <b>local</b> NEO blockchain with a <b>single</b> command.
</p>

## What?

- Writing smart contracts requires a local version of the NEO blockchain to be running.
- There are multiple services required to get this working.
- This repo uses [Docker Compose](https://docs.docker.com/compose/) to package and run all of the services.

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

## Components

The [Docker Compose](https://docs.docker.com/compose/) stack is made up of the following 
services:

- [neo-privatenet](https://hub.docker.com/r/cityofzion/neo-privatenet/) (consensus nodes)
- [neo-python](https://github.com/CityOfZion/neo-python) (development CLI)
- [neo-scan](https://github.com/CityOfZion/neo-scan) (block explorer)
- [postgres](https://hub.docker.com/_/postgres/) (database for neo-scan)

## Credit

This is a fork of [@slipo](https://github.com/slipo)'s 
[neo-scan-docker](https://github.com/slipo/neo-scan-docker) repo!

