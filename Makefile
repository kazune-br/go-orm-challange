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
	sql-migrate new ${FILE} -config=$$(pwd)/db/gorm/dbconfig.yml

migration:
	sql-migrate up -config=./db/gorm/dbconfig.yml
	sql-migrate up -config=./db/sqlboiler/dbconfig.yml

status:
	sql-migrate status -config=./db/gorm/dbconfig.yml
	sql-migrate status -config=./db/sqlboiler/dbconfig.yml

gormdb:
	docker-compose $(COMPOSE_OPTS) exec gormdb mysql -uroot -ppassword -Dgormdb

sqlboilerdb:
	docker-compose $(COMPOSE_OPTS) exec sqlboilerdb mysql -uroot -ppassword -Dsqlboilerdb

delete-volume:
	rm -rf ./db/mysql_data
