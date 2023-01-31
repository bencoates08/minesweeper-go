# For testing full build and deployment
development:
	docker-compose -f docker/dev/docker-compose.yml up -d

down-development:
	docker-compose -f docker/dev/docker-compose.yml down

# For testing local development
local:
	docker-compose -f docker/local-dev/docker-compose.yml up -d

down-local:
	docker-compose -f docker/local-dev/docker-compose.yml down

run-frontend:
	yarn --cwd frontend dev

dev:
	make local && make run-frontend

down:
	make down-local