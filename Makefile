.PHONY: setup build up down logs test db create-test-db

BASE_DOCKER_COMPOSE = docker-compose.yml
COMPOSE_OPTS        = -f "$(BASE_DOCKER_COMPOSE)"
FILE				= create_sample

setup: migration migration-test
clean: down delete-volume

up:
	docker-compose $(COMPOSE_OPTS) up -d

down:
	docker-compose $(COMPOSE_OPTS) down

logs:
	docker-compose $(COMPOSE_OPTS) logs -f

new-migration:
	sql-migrate new ${FILE} -config=./db/gorm/dbconfig.yml

migration:
	sql-migrate up -config=./db/gorm/dbconfig.yml
	sql-migrate up -config=./db/sqlboiler/dbconfig.yml

gorm:
	docker-compose $(COMPOSE_OPTS) exec gormdb mysql -uroot -ppassword -Dgormdb

delete-volume:
	rm -rf ./db/mysql_data
