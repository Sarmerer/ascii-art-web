docker-run:
	docker run --name ascii-art-web -p 4435:4435 -d ascii-art-web
docker-build:
	docker build -f Dockerfile -t ascii-art-web .
dockerize: docker-build docker-run
go-build:
	bash -c "go build -o main"
go:
	bash -c "go run main.go"