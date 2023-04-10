## Build and start docker compose
up: 
	@echo "Building and starting docker compose"
	docker-compose up -d --build
	@echo "Docker compose started"

## Stop and remove docker compose
down: 
	@echo "Stopping and removing docker compose"
	docker-compose down
	@echo "Docker compose stopped"

## deploy auth cloudformation stack
deploy-auth-cf:
	@echo "Deploying auth cloudformation stack"
	aws cloudformation deploy --template-file ../cloudformation/auth.yml --stack-name auth --capabilities CAPABILITY_NAMED_IAM
	@echo "Auth cloudformation stack deployed"

## deploy notes cloudformation stack
deploy-notes-cf:
	@echo "Deploying notes cloudformation stack"
	aws cloudformation deploy --template-file ../cloudformation/notes.yml --stack-name notes --capabilities CAPABILITY_NAMED_IAM
	@echo "Notes cloudformation stack deployed"

## deploy k8s cloudformation stack
deploy-k8s-cf:
	@echo "Deploying k8s cloudformation stack"
	aws cloudformation deploy --template-file ../cloudformation/k8s.yml --stack-name k8s --capabilities CAPABILITY_NAMED_IAM
	@echo "K8s cloudformation stack deployed"

## down and up docker compose
restart: deploy-auth-cf deploy-notes-cf down up