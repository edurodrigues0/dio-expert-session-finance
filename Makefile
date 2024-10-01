build-image:
	docker build -t edurodrigues0/finance .

run-app:
	docker-compose -f .devops/app.yml up -d

unit_test:
	go test ./...