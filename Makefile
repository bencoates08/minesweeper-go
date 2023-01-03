run_docker:
	docker-compose -f docker/docker-compose.yml up -d

stop_docker:
	docker-compose -f docker/docker-compose.yml down

run_app:
	yarn --cwd frontend dev

dev:
	make run_docker && make run_app

kill:
	make stop_docker