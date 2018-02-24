DEFAULT: start

start:
	@./scripts/print.sh "Starting Docker containers..."
	@docker-compose up -d > /dev/null
	@./scripts/print.sh "Waiting for network..." false
	@./scripts/ping.sh
	@./scripts/print.sh "Done 🎉"

start-verbose:
	@./scripts/print.sh "Starting Docker containers..."
	docker-compose up -d
	@./scripts/print.sh "Waiting for network..." false
	@./scripts/ping.sh
	@./scripts/print.sh "Done 🎉"

stop:
	@./scripts/print.sh "Stopping Docker containers..."
	@docker-compose down > /dev/null
	@./scripts/print.sh "Done 🎉"