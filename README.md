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
- Ensure all developers are on same env
- Run each component in a separate container with its own dependencies and libraries
- Build Once, Use Many by developers
- Docker (high level) utilizes Aleksey (low level) containers
- Sharing the kernel
- Ship, run anywhere, anytime, many times
- Efficiency than VM
- Easy recovery
- Enabler of DevOps

## Prerequisite 

- to follow along this workshop you need to have docker engine ready running in your system
- I recommend to use Linux native system ether your local system or AWS EC2 Linux, to experience 100% of docker features
- If you are using Docker Desktop on Windows/MacOSX, it will be not native, instead the Docker Desktop will spine up VM on top of your OS
- verify your docker engine ready by `version` command

```shell
docker version
```

output:

![docker-version](screenshots/docker-version.png)

## Practise Commands and capture screen:

- run : start a container

```shell
docker run $image_name # attach mode, CTRL+C to stop, container will also exit
```

![docker-run-attach-mode](screenshots/docker-run-attach-mode.png)

```shell
docker run -d $image_name # detach mode
```

![docker-run-detach-mode](screenshots/docker-run-detach-mode.png)

```shell
docker run -i $image_name # interactive mode
```

![docker-run-interactive-mode](screenshots/docker-run-interactive-mode.png)

```shell
docker run -it $image_name # interactive sudo terminal mode
```

![docker-run-it](screenshots/docker-run-it.png)

```shell
docker run -p $external_port:$internal_port $image_name # mapping port 8080 of localhost to port 80 of container, we can't now access nginx through port 80, but instead need to use 8080.
```

![docker-run-port-mapping-01](screenshots/docker-run-port-mapping-01.png)
![docker-run-port-mapping-02](screenshots/docker-run-port-mapping-02.png)

```shell
docker run -v $docker_host_dir:$container_dir $image_name # persistent data store
```

![volume-not-work](screenshots/volume-not-work.png)
for Docker Desktop in Windows/MacOSX will need to search for work around on docker volume.

- attach : attach terminal into running container

```shell
docker attach $container_name
docker attach $container_id
```

![docker-attach](screenshots/docker-attach.png)

when we attach without interactive mode, the container will stop if we do `CTRL+C`
but if just close our terminal window, will not impact to attached console, so container still running.

- ps : list containers

```shell
docker ps # show only running containers
```

![docker-ps](screenshots/docker-ps.png)

```shell
docker ps -a # include stopped containers
```

![docker-ps-a](screenshots/docker-ps-a.png)

- stop : stop a container

```shell
docker stop $container_id # use container id as argument
docker stop $container_name # use container name as argument
```

![docker-stop](screenshots/docker-stop.png)

- rm : remove an exited container

```shell
docker rm $container_id # use container id as argument
docker rm $container_name # use container name as argument
```

![docker-rm](screenshots/docker-rm.png)

For exploration:

- remove multiple stopped containers in single command
- remove running container

- images : list images

```shell
docker images # deprecated
docker image ls # new version
```

![docker-image-ls](screenshots/docker-image-ls.png)

container life cycle: image -> container -> stop -> remove

- rmi : remove image

```shell
docker rmi $image_name
docker image rm $image_name
docker rmi $image_id
docker image rm $image_id
```

![docker-image-rm](screenshots/docker-image-rm.png)

For exploration:

- try to remove image that running by a container

- pull : download image

```shell
docker pull $repo_name
```

![docker-pull](screenshots/docker-pull.png)

- tag : create new version of image

```shell
docker tag $existing_image_full_name $new_image_full_name
```

![docker-tag](screenshots/docker-tag.png)

- push : upload image, require to have docker hub account

```shell
docker push $docker_hub_account/$docker_image_name:$tag
```

![docker-push-01](screenshots/docker-push-01.png)

![docker-push-02](screenshots/docker-push-02.png)

- exec : execute a command in running container

```shell
docker exec $container_name <command>
```

![docker-exec](screenshots/docker-exec.png)

- logs : view the logs inside the container

```shell
docker logs $container_name
docker logs $container_id
```

![docker-logs](screenshots/docker-logs.png)
