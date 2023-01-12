all:
	$(MAKE) -C app all
	$(MAKE) -C web-client all

init-dev-db:
	$(MAKE) -C app init-dev-db

run-dev-server:
	$(MAKE) -C app run-dev

run-dev-client:
	$(MAKE) -C web-client run-dev