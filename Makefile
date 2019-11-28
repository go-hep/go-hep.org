.PHONY: all build deploy clean

HUGO ?= hugo

build: clean
	$(HUGO)

deploy: build
	go run cmd_deploy.go

serve:
	$(HUGO) serve -w -p 8080

clean:
	/bin/rm -rf public
