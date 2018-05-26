BRANCH ?= "master"
VERSION ?= $(shell cat ./VERSION)

DEFAULT: start

# IGNORE - used to check if version has been bumped on CI.
check-version:
	@echo "=> Checking if VERSION exists as Git tag..."
	(! git rev-list ${VERSION})

# IGNORE - used to test project on CI.
integration-tests: setup-network
	@./scripts/print.sh prefix "Checking number of running services"
	@./scripts/status.sh

# IGNORE - used to release new version of project to Git.
push-tag:
	@echo "=> New tag version: ${VERSION}"
	git checkout ${BRANCH}
	git pull origin ${BRANCH}
	git tag ${VERSION}
	git push origin ${BRANCH} --tags

setup-network:
	@./scripts/print.sh prefix "Fetching Docker containers..."
	@docker-compose pull > /dev/null
	@./scripts/print.sh prefix "Starting Docker containers..."
	@docker-compose up -d --build --remove-orphans --force-recreate > /dev/null
	@./scripts/print.sh prefix "Waiting for network..." false
	@./scripts/ping.sh
	@./scripts/print.sh prefix "Network running! 🎉"

start: setup-network
	@./scripts/print.sh prefix "Attaching terminal to neo-python client\n"
	@./scripts/print.sh grey "Open wallet (password: 'coz'):\t open wallet ./neo-privnet.wallet"
	@./scripts/print.sh grey "Test smart contract:\t\t build /smart-contracts/wake_up_neo.py test 07 05 True False main\n"
	@docker exec -it neo-python np-prompt -p -v

stop:
	@./scripts/print.sh prefix "Stopping Docker containers..."
	@docker-compose down > /dev/null
	@./scripts/print.sh prefix "Done 🎉"
