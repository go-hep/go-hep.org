.PHONY: all build deploy clean

build: clean
	hugo

deploy: build
	go run cmd_deploy.go

serve:
	hugo serve -w

clean:
	/bin/rm -rf public
