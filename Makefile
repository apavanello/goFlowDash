DEV:
	cd assets/src && yarn vite build --outDir ../../assets/dist && cd ../..
	go run cmd/main.go

DBA:
	docker-compose -f docker-compose.yml up -d