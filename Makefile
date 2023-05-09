up:
	docker-compose -f ./.infrastructure/docker-compose.yml up --force-recreate --build -d

down:
	docker-compose -f ./.infrastructure/docker-compose.yml down

lint:
	golangci-lint run -v --timeout 15m --fix cmd/... internal/...