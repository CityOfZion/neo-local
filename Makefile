DEFAULT: start

start:
	@./scripts/print.sh prefix "Starting Docker containers..."
	@docker-compose up -d --build --remove-orphans --force-recreate > /dev/null
	@./scripts/print.sh prefix "Waiting for network..." false
	@./scripts/ping.sh
	@./scripts/print.sh prefix "Network running!"
	@./scripts/print.sh prefix "Attaching terminal to neo-python client\n"
	@./scripts/print.sh grey "Enable logging:\t\t\t config sc-events on"
	@./scripts/print.sh grey "Open wallet (password: 'coz'):\t open wallet ./neo-privnet.wallet\n"
	@docker exec -it neo-python python prompt.py -p

stop:
	@./scripts/print.sh "Stopping Docker containers..."
	@docker-compose down > /dev/null
	@./scripts/print.sh "Done 🎉"
