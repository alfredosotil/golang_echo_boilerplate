lint_docker_compose_file = "./development/golangci_lint/docker-compose.yml"

lint-build:
	@echo "🌀 ️container are building..."
	@docker-compose --file=$(lint_docker_compose_file) build -q
	@echo "✔  ️container built"

lint-check:
	@echo "🌀️ code linting..."
	@docker-compose --file=$(lint_docker_compose_file) run --rm echo-golinter golangci-lint run \
 		&& echo "✔️  checked without errors" \
 		|| echo "☢️  code style issues found"

lint-fix:
	@echo "🌀 ️code fixing..."
	@docker-compose --file=$(lint_docker_compose_file) run --rm echo-golinter golangci-lint run --fix \
		&& echo "✔️  fixed without errors" \
		|| (echo "⚠️️  you need to fix above issues manually" && exit 1)
	@echo "⚠️️ run \"make lint-check\" again to check what did not fix yet"

compose-up: ## Starts a local instance of the service with instances in docker
	@if [ -f docker-compose.yml ]; then\
		docker-compose -p ws -f docker-compose.yml up -d ;\
	else\
		echo "No compose file.";\
	fi

compose-down: ## Shutdowns the local containers
	@docker-compose -p ws -f docker-compose.yml down --rmi local