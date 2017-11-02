# How to setup

To build the container image

```
docker build -t klim-feeder-app .
```

To run the container

```
docker run -it --rm -p 5672 --network="develk_elk"  --name my-running-app geertvl/klim-feeder-app:2.0
```