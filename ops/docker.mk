# Docker
up:
	@docker-compose up -d

stop:
	@docker-compose stop

down:
	@docker-compose down

logs:
	@docker-compose logs -f

enter:
	@docker exec -it grpc_$(service)_1 /bin/sh