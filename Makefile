run:
	docker-compose up

down:
	docker-compose down

migrate:
	docker-compose run --rm migrate

drop-tables:
	docker exec -i web3appdemo-db psql -U postgres -d web3appdemo -c \
	"DROP SCHEMA public CASCADE; CREATE SCHEMA public;"

gen-model:
	sqlboiler psql --no-tests

