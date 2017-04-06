.PHONY: all build deploy clean

build: clean
	hugo

deploy: build
	go run cmd_deploy.go

serve:
	hugo serve -w -p 8080

clean:
	/bin/rm -rf public
