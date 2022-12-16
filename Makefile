include .env

migrate-up :
	migrate -source file:database/migrations -database postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable up $(N)
migrate-down :
	migrate -source file:database/migrations -database postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable down $(N)
migrate-force :
	migrate -source file:database/migrations -database postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable force $(N)