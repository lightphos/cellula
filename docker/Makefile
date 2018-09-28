default: up

up: 
	docker rm cellx || true
	docker run --name cellx java:8 true
	docker-compose rm -f
	docker-compose build
	docker-compose up 
