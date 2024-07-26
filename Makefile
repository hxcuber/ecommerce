PROJECT_NAME := ecommerce
COMPOSE := PROJECT_NAME=${PROJECT_NAME} docker compose -f deploy/compose.yaml -p ${PROJECT_NAME}

docker-build:
	docker build -f build/api.Dockerfile -t ${PROJECT_NAME}-api:latest .
	docker build -f build/migrate.Dockerfile -t ${PROJECT_NAME}-migrate:latest .
	docker image prune -f

docker-setup:
	${COMPOSE} up pg -d
	sleep 5
	${COMPOSE} up db-migrate
	@echo 'db set up and migrated, running in background'

docker-run:
	@echo 'db is running in background, "make docker-stop-db" to stop db'
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

docker-boilerplate:
	docker start-db
	PSQL_DBNAME=${PROJECT_NAME} PSQL_HOST=pg PSQL_USER=${PROJECT_NAME} docker compose run 'sh -c sqlboiler --wipe --add-enum-types psql'
	docker stop-db

local-boilerplate:
	PSQL_DBNAME=${PROJECT_NAME} PSQL_HOST=localhost PSQL_USER=hxcuber sh -c 'sqlboiler --wipe --add-enum-types psql'

