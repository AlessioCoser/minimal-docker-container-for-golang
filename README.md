# Minimal Docker Container for GoLang
The image starts from `scratch`: a special docker image that’s empty. It’s truly 0B!

A normal compiled app is dynamically linked to the libraries it needs to run (i.e., all the C libraries it binds to).

Unfortunately, scratch is empty, so there are no libraries and no loadpath for it to look in. What we have to do is modify our build script to statically compile our app with all libraries built in.

## 1. Build go executable for linux
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .
```

## 2. Build Docker image
```
docker build -t go-docker-minimal .
```

This docker image is less than 2 Mb!
```
REPOSITORY                        TAG                  IMAGE ID            CREATED             SIZE
go-docker-minimal                 latest               a5ad46a94d25        5 minutes ago       1.88MB
```

## 3. Run Docker container
```
docker run -t go-docker-minimal
```

## Credits
To have more information about it [read this article from codeship.com](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/)