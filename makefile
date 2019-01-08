assets-build:
	rm -rf config/assets.go
	go-assets-builder config/local.yml assets/templates/* > config/assets.go -p config