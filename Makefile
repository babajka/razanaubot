HOST = "wir-prod@wir.by"

run:
	cp secrets/local.yaml pkg/config/secret/secret.yaml
	go run cmd/bot/main.go
	rm pkg/config/secret/secret.yaml

deploy-production:
	cp secrets/production.yaml pkg/config/secret/secret.yaml
	GOOS=linux GOARCH=amd64 go build -o build/razanaubot ./cmd/bot
	scp ./build/razanaubot $(HOST):/home/wir-prod/deployed/
