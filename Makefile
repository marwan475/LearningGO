

migrationuser:
	migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users

migrationposts:
	migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_posts


migrationalterposts:
	migrate create -seq -ext sql -dir ./cmd/migrate/migrations alter_posts

migrateup:
	migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/testdb?sslmode=disable" up

migratedown:
	migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/testdb?sslmode=disable" down

