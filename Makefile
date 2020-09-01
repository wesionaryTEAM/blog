firebase:
	$(MAKE) -C "services/firebaseLogin/"

serve: 
	fresh 
dev:
	go run main.go