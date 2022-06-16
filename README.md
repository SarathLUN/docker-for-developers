# `</>` Docker for Developers `</>`;

#### Date: 2022-06-17;
#### Time: 14:00 - 15:00;
#### Location: CADT;

---

## Why do you need docker?

- OS compatibility
- Library/dependency compatibility
- Resolve [The Matrix from Hell](https://dzone.com/articles/are-you-stuck-in-the-new-devops-matrix-from-hell) !!!
- New developer joiner, setup environment
- Ensure all developers are on same env
- Run each component in a separate container with its own dependencies and libraries
- Build Once, Use Many by developers
- Docker (high level) utilizes Aleksey (low level) containers
- Sharing the kernel
- Ship, run anywhere, anytime, many times
- Efficiency than VM
- Easy recovery
- One Enabler for DevOps adaption

## Commands

- help: show all available functions of the docker command

```shell
docker --help 
docker COMMAND --help
docker [OPTIONS] COMMAND
```

- run : start a container

```shell
docker run $image_name # attach mode, CTRL+C to stop, container will also exit
docker run -d $image_name # detach mode
docker run -i $image_name # interactive mode
docker run -it $image_name # interactive sudo terminal mode
docker run -p $external_port:$internal_port $image_name # mapping port 8080 of localhost to port 80 of container, we can't now access nginx through port 80, but instead need to use 8080.
docker run -v $docker_host_dir:$container_dir $image_name # persistent data store
```

- attach : attach terminal into running container

```shell
docker attach $container_name
docker attach $container_id
```

- ps : list containers

```shell
docker ps
docker ps -a # include stopped containers
```

- stop : stop a container

```shell
docker stop $container_id # use container id as argument
docker stop $container_name # use container name as argument
```

- rm : remove an exited container

```shell
docker rm $container_id # use container id as argument
docker rm $container_name # use container name as argument
```

- images : list images

```shell
docker images # deprecated
docker image ls # new version
```

- rmi : remove image

```shell
docker rmi $image_name
docker image rm $image_name
docker rmi $image_id
docker image rm $image_id
```

- pull : download image

```shell
docker pull $repo_name
```

- push : upload image

```shell
docker push $docker_hub_account/$docker_image_name:$tag
```

- exec : execute a command in running container

```shell
docker exec $container_name <command>
```

- logs : view the logs inside the container

```shell
docker logs $container_name
docker logs $container_id
```

## Environment Variables

- set OS env via `run` command

```shell
docker run -e $env_name=$env_value $image_name
```

- show OS env via `inspect` command

```shell
docker inspect $container_name
```

## How to create my own image?

- step 1: identify the steps to be done in the right order:

1. OS - golang:alpine
2. define working directory
3. copy source code into container
4. initialize go module
5. download application dependencies
6. build our application into binary executable file
7. run our program

- step 2: convert the steps into `Dockerfile`

```dockerfile
FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN go mod init
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]

```

- step 3: build image

```shell
# docker build -t $docker_hub_account/$image_name:$tag .
docker build -t tonysarath/my-golang-app:latest .

# by default "Dockerfile" will be used, but we can also specify it
# Ex: docker build -t $docker_hub_account/$image_name:$tag $path_to_dockerfile
```

- step 4: run container

```shell
docker run -itd -p 80:80 --rm --name my-running-app tonysarath/my-golang-app
```