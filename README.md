# Introduction

This contains Docker files and related configuration files for a container of [neo-scan](https://github.com/CityOfZion/neo-scan).

`Dockerfile` is what is used to build [the docker container](https://hub.docker.com/r/slipoh/neo-scan/) found on Docker Hub.

`docker-compose.yml` is a a fully working private network with neo-scan connected to it. It uses prebuilt images.

# Privnet with GAS and with neo-scan — quick start

Note: This is a complete solution. It expects you're not already running the privnet docker image. It will set up a new one. You may need to remove your currently running docker privnet image if you already have one.

It uses the prebuilt [neo-privnet-with-gas](https://hub.docker.com/r/metachris/neo-privnet-with-gas/]) and [neo-scan](https://hub.docker.com/r/slipoh/neo-scan/) images.

We assuming you have Docker all set up.

Pull the repo down and `cd` into it
```
git clone git@github.com:slipo/neo-scan-docker.git
cd neo-scan-docker
```

Start up the container
```
docker-compose up
```

It will take some time to set up.

While you wait, add this line to your hosts file:
```
127.0.0.1 neo-privnet
```

That allows you to connect to the privnet NEO nodes with the URLs returned by the neo-scan container. If you're using neo-python to connect to the privnet, you can use the standard configuration. 127.0.0.1:30333 will continue to work, for example.

OK, if you waited a few minutes, it should be ready. Check: http://localhost:4000/. You should see neo-scan with some blocks.

For information about the privnet setup, including the keys to your assets check [the neo-privnet-with-gas documentation](https://hub.docker.com/r/metachris/neo-privnet-with-gas/)
