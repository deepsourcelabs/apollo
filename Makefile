swagger:
	swag init -g cmd/apollo/routes/routes.go

local_build:
	cd cmd/apollo && GOOS=linux GOARCH=amd64 go build -tags static_all -o /tmp/apollo .

local_run:
	/tmp/apollo config --path config.yaml

dev:
	make local_build
	make local_run
