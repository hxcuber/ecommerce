PROJECT_NAME := ecommerce
COMPOSE := PROJECT_NAME=${PROJECT_NAME} docker compose -f deploy/compose.yaml -p ${PROJECT_NAME}

docker-build:
	docker build -f build/api.Dockerfile -t ${PROJECT_NAME}-api:latest .
	docker build -f build/migrate.Dockerfile -t ${PROJECT_NAME}-migrate:latest .
	docker image prune -f

docker-run:
	${COMPOSE} up pg -d
	sleep 5
	${COMPOSE} up db-migrate
	${COMPOSE} up api

docker-start-db:
	${COMPOSE} start pg
docker-stop-db:
	${COMPOSE} stop pg

docker-teardown:
	${COMPOSE} down
	docker image prune -f

docker-test:
	${COMPOSE} up test
	${COMPOSE} down test