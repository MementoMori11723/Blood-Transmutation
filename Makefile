run :
	@docker-compose -f config/compose.yml -p blood up --build -d

stop :
	@docker-compose -f config/compose.yml -p blood down --remove-orphans
