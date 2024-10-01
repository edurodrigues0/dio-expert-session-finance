build-image:
	docker build -t edurodrigues0/finance .

push-image:
	docker push marcopollivier/finance

run-app:
	docker-compose -f .devops/app.yml up -d

stop-app:
	docker-compose -f .devops/app.yml down

prepare-env:
	docker-compose -f .devops/postgres.yml up -d

unit_test:
	go test ./...