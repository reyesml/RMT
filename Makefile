all:
	$(MAKE) -C app all
	$(MAKE) -C web-client all

run-dev-server:
	$(MAKE) -C app run-dev

run-dev-client:
	$(MAKE) -C web-client run-dev