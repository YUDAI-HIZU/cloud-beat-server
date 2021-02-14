.PHONY: shell
shell:
	docker exec -it cloud-beat-server_app_1 sh

.PHONY: run
run:
	docker-compose run --rm app

.PHONY: gen
gen:
	docker-compose run --rm app go run github.com/99designs/gqlgen
