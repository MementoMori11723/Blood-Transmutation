run :
	@docker-compose -f compose.yml -p blood up --build -d

stop :
	@docker-compose -f compose.yml -p blood down --remove-orphans
