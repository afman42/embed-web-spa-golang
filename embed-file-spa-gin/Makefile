build:
	@echo "Mod Tidy.."
	go mod tidy;
	@echo 'Building...'
	GOOS=linux GOARCH=amd64 go build -ldflags="-s" -o=./embed-web-spa-gin .
	@echo 'Finish Building'

deploy/dockerup:
	@echo 'Deploying...'
	sudo docker build -t go-embed-spa-prod -f Dockerfile.production .;
	sudo docker run -dit --rm -p 3030:3030 --network=app-net --name="goembedspaprod" --restart=always go-embed-spa-prod;
	@echo 'Finish Deploying'

deploy/dockercomposeup:
	@echo 'Deploying...'
	docker compose up -d go-embed-spa-prod
	@echo 'Finish Deploying'