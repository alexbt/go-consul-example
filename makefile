export MYHOST = $(shell ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1' | awk '{print $1}')

compose:
	pushd go-service-one && make docker-build; popd
	pushd go-service-two && make docker-build; popd
	docker-compose up
