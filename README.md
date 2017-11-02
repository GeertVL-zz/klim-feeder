# How to setup

To build the container image

```
docker build -t klim-feeder-app .
```


To run the container

```
docker run -it --rm -p 5672 --network="develk_elk"  --name my-running-app geertvl/klim-feeder-app:2.0
```


Tag the container for upload

```
docker tag klim-feeder-app $DOCKER_ID_USER/klim-feeder-app:2.0
```

ps.: 2.0 is the tag 


Pushing to the docker cloud

```
docker login (uid/pwd needed)
docker push $DOCKER_ID_USER/klim-feeder-app
```
