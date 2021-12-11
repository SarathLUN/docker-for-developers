# `</>` Docker for Developers `</>`;

#### WORKSHOP # 02;
#### Date: 2021-12-12;
#### Time: 10:00 - 16:00;
#### Location: Microsoft Teams;

---

## Why do you need docker?

- OS compatibility
- Library/dependency compatibility
- The Matrix from Hell!!
- New developer joiner, setup environment
- Make sure all developers are on same env
- Run each component in a separate container with its own dependencies and libraries
- Build Once, Use Many by developers
- Docker (high level) utilizes Aleksey (low level) containers
- Sharing the kernel
- Ship, run anywhere, anytime, many times
- Efficiency than VM
- Easy recovery
- Enabler of DevOps

## Commands

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

- identify the steps to be done in the right order:

  1. OS - Ubuntu
  2. Update apt repo
  3. Install dependencies using `apt`
  4. Install python dependencies using `pip`
  5. Copy source code to `/opt` folder
  6. Run the web server using `flask` command

- convert the steps into `Dockerfile`

```dockerfile
FROM Ubuntu

RUN apt-get update
RUN apt-get install python

RUN pip install flask
RUN pip install flask-mysql

COPY . /opt/source-code

ENTRYPOINT FLASK_APP=/opt/source-code/app.py flask run
```

- build image

```shell
docker build -t $docker_hub_account/$image_name:$tag .

# by default "Dockerfile" will be used, but we can also specify it
docker build -t $docker_hub_account/$image_name:$tag $path_to_dockerfile
```
