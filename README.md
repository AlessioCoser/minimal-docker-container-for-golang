# Minimal Docker Container for GoLang
The image starts from `scratch`: a special docker image that’s empty. It’s truly 0B!

A normal compiled app is dynamically linked to the libraries it needs to run (i.e., all the C libraries it binds to).

Unfortunately, scratch is empty, so there are no libraries and no loadpath for it to look in. What we have to do is modify our build script to statically compile our app with all libraries built in.

## Build and Run an Hello World app
### 0. Prerequisites
- [Docker](https://docs.docker.com/install)
- [golang](https://golang.org/doc/install)

### 1. Build go executable for linux
```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .
```

### 2. Build Docker image
```sh
docker build -t go-docker-minimal .
```

The docker image created is less than 2 Mb. Running `docker images` you can see this:
```
REPOSITORY                        TAG                  IMAGE ID            CREATED             SIZE
go-docker-minimal                 latest               a5ad46a94d25        5 minutes ago       1.88MB
```

### 3. Run Docker container
```sh
docker run -t go-docker-minimal
```

## Build and Execute with a multi-stage build
Using a multi-stage build you can compile go executable and then create the minimal docker-image in the same build instruction.

### Build
(With `-f` I can pass a non default dockerfile filename)
```
docker build -t go-docker-minimal-multistage -f Dockerfile.multistage .
```

### Run
```sh
docker run -t go-docker-minimal-multistage
```

## Credits
To have more information about it [read this article from codeship.com](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/)