# Minimal Docker Container for GoLang

## 1. Build go executable for linux
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .
```

## 2. Build Docker image
```
docker build -t go-docker-minimal .
```

## 3. Run Docker container
```
docker run -t go-docker-minimal
```
