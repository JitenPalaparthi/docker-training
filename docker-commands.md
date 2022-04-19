### To check whether docker is installed

```docker version```

```docker info```

### To pull an ubuntu image

```docker pull ubuntu```

### To check images

```docker image ls```

```docker images```

### To create a container

```docker run -i -t --name u1 ubuntu bash```

### To create a container on detached mode and publish a port

```docker run -d --name n1 -p 50080:80 nginx```

### To list running containers

```docker container ls```

```docker ps```

### To list all containers running and stopped

```docker ps -a```

### To remove a container 

```docker rm u1```

### To remove stopped containers

```docker container prune```

### To forcely remove a running container

```docker rm -f u1```

### To remove an image

```docker rmi ff0fea8310f3```

### To remove unused images

```docker image prune```

### To remove an image forcefully

```docker rmi -f ff0fea8310f3```

### To fetch only imageids

```docker images -q```

### To remove all images 

```docker rmi -f $(docker images -q)```

### To remove all containers

```docker rm -f $(docker ps -aq)```

### To stop a container

```docker stop u1```

### To start a stopped container

```docker start u1```

### To access running container

```docker exec -it u1 bash```

### To get logs of a container

```docker logs u1```

### Create a new network

```docker network create mybridge```

### Create a new network with gateway

```docker network create --gateway 172.28.0.1 --subnet 172.28.0.0.16```

### Copy a file from host to container

```docker cp tempfile n3:.```

### Copy a file from container to host

```docker cp n3:docker-entrypoint.sh .```

### To create a volume

```docker volume create myvolume```

### To use the volume for a container

```docker run -d --name n3 --volume myvolume:/word nginx```

### To mount host directory to contaier directory

```docker run -d -it --name n2 --mount type=bind,source="$(pwd)"/workspace,target=/myworkspace nginx```

### To mount host directoy using -v option

```docker run -d --name n4 -v "$(pwd)"/workspace:/myworkspace nginx```

### Container registries

Docker hub
acr
gcr
quay

### Dockerfiles

- Docker file consistrs of few commands
- The default file name is Dockerfile. But can give any file name.
- There should be a base image

- Create an ubuntu image with vim installed.
- the manual steps that are involved are
```apt-get update -y```
```apt-get upgrade -y```
```apt-get install vim -y```

- RUN command is used to execute commands in the shell
- Every RUN command creates a new layer

### How to build docker file

```docker build . -t jpalaparthi/ubuntuplus```

### How to push the image to the registry

```docker push jpalaparthi/ubuntuplus```

### Build docker file with another filename

```docker build . -f PingDockerfile -t jpalaparthi/pingapp```