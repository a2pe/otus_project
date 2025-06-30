.PHONY: run build down logs

build:
	docker-compose build

run:
	docker-compose up

down:
	docker-compose down

logs:
	docker-compose logs -f app

test:
	docker exec -it go_app go test ./...