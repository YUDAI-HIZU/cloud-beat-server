.PHONY: shell
shell:
	docker exec -it cloud-track-backend_app_1 sh

.PHONY: run
run:
	docker-compose run --rm app
