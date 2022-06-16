# My Solution:
```shell
>docker build -t my-golang-app .
[+] Building 3.0s (11/11) FINISHED
 => [internal] load build definition from Dockerfile                                                                                                                                0.0s
 => => transferring dockerfile: 168B                                                                                                                                                0.0s
 => [internal] load .dockerignore                                                                                                                                                   0.0s
 => => transferring context: 34B                                                                                                                                                    0.0s
 => [internal] load metadata for docker.io/library/golang:latest                                                                                                                    1.6s
 => [1/6] FROM docker.io/library/golang:latest@sha256:6f0b0a314b158ff6caf8f12d7f6f3a966500ec6afb533e986eca7375e2f7560f                                                              0.0s
 => [internal] load build context                                                                                                                                                   0.0s
 => => transferring context: 226B                                                                                                                                                   0.0s
 => CACHED [2/6] WORKDIR /go/src/app                                                                                                                                                0.0s
 => [3/6] COPY . .                                                                                                                                                                  0.0s
 => [4/6] RUN go mod init                                                                                                                                                           0.2s
 => [5/6] RUN go get -d -v ./...                                                                                                                                                    0.3s
 => [6/6] RUN go install -v ./...                                                                                                                                                   0.6s
 => exporting to image                                                                                                                                                              0.1s
 => => exporting layers                                                                                                                                                             0.1s
 => => writing image sha256:0880b7eff77d2d545478859d4568f352bb93b5307b4e8e906f55583eb8fe1fa1                                                                                        0.0s
 => => naming to docker.io/library/my-golang-app                                                                                                                                    0.0s

Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
```

```shell
>docker images
REPOSITORY        TAG           IMAGE ID       CREATED             SIZE
my-golang-app     latest        0880b7eff77d   18 seconds ago      868MB
my-curl           latest        93bd0b0da497   About an hour ago   117MB
my-docker-whale   latest        3698a88acc62   2 days ago          278MB
my-docker-whale   entry-point   5351b4ddde26   2 days ago          278MB
hello-go          latest        86770b7e7759   9 days ago          862MB
docker/whalesay   latest        6b362a9f73eb   5 years ago         247MB
```

```shell
>docker run -itd -p 80:80 --rm --name my-running-app my-golang-app
325866d4181cf12a973703db5f86c1ffc52bdf42d99bba8302a798543664f192
```
- check running processes:
```shell
>docker ps
CONTAINER ID   IMAGE           COMMAND   CREATED         STATUS         PORTS                               NAMES
77eefc4204b8   my-golang-app   "app"     4 seconds ago   Up 3 seconds   0.0.0.0:80->80/tcp, :::80->80/tcp   my-running-app
```

- Note: `--rm` will auto remove container after we stop it
```shell
>telnet localhost 80
Trying ::1...
Connected to localhost.
Escape character is '^]'.

```

```shell
>curl http://localhost/
Hello from docker container!
```