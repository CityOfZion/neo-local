BRANCH ?= "master"
VERSION ?= $(shell cat ./VERSION)

DEFAULT: start

attach-to-neo-python-client:
	@./scripts/print.sh prefix "Attaching terminal to neo-python client\n"
	@./scripts/print.sh grey "Open wallet (password: 'coz'):\t open wallet ./neo-privnet.wallet"
	@./scripts/print.sh grey "Test smart contract:\t\t build /smart-contracts/wake_up_neo.py test 07 05 True False main\n"
	@docker exec -it neo-python np-prompt -p -v

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

pull-images:
	@./scripts/print.sh prefix "Fetching Docker containers..."
	@docker-compose pull > /dev/null

setup-network:
	@./scripts/print.sh prefix "Starting Docker containers..."
	@./scripts/print.sh prefix "The first time you run, it will take a while to build neo-cli-privatenet (1 to 4) images..."
	@docker-compose up -d --build --force-recreate --remove-orphans > /dev/null
	@./scripts/print.sh prefix "Waiting for network..." false
	@./scripts/ping.sh
	@./scripts/print.sh prefix "Network running! 🎉"

start: pull-images setup-network attach-to-neo-python-client

start-offline: setup-network attach-to-neo-python-client

stop:
	@./scripts/print.sh prefix "Stopping Docker containers..."
	@docker-compose down > /dev/null
	@./scripts/print.sh prefix "Done 🎉"
