# Minimal Docker Container for GoLang
The image starts from `scratch`: a special docker image that’s empty. It’s truly 0B!

A normal compiled app is dynamically linked to the libraries it needs to run (i.e., all the C libraries it binds to).

Unfortunately, scratch is empty, so there are no libraries and no loadpath for it to look in. What we have to do is modify our build script to statically compile our app with all libraries built in.

# Build and Run an Hello World app

## First Way
Build on your machine

#### 0. Prerequisites
- [Docker](https://docs.docker.com/install)
- [golang](https://golang.org/doc/install)

#### 1. Build go executable for linux
```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .
```

Alternatively you can use the `-s` and `-w` linker flags to strip the debugging information and reduce container size
```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o main .
```

#### 2. Build Docker image
```sh
docker build -t go-docker-minimal-api .
```

#### 3. Run Docker container
```sh
docker run -p 8080:8080 -t go-docker-minimal-api
```

## Second Way
Using a multi-stage build you can compile go executable and then create the minimal docker-image using the same dockerfile.

#### 0. Prerequisites
- [Docker](https://docs.docker.com/install)

#### 1. Build
(With `-f` I can pass a non default dockerfile filename)
```
docker build -t go-docker-minimal-api-multistage -f Dockerfile.multistage .
```

#### 2. Run
```sh
docker run -p 8080:8080 -t go-docker-minimal-api-multistage
```

#### 3. Try
```sh
curl localhost:8080/helloworld
# => {"greeting":"Hello World"}
```

## How much minimal?
The docker image created is less than 5 Mb. Running `docker images` you can see this:
```
REPOSITORY                        TAG                  IMAGE ID            CREATED             SIZE
go-docker-minimal-api             latest               a5ad46a94d25        5 minutes ago       4.16MB
```

# Credits
To have more information about it [read this article from codeship.com](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/)
