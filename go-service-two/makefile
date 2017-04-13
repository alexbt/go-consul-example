.VERSION:=1.0.0
.BRANCH:=`git branch | grep \* | cut -d ' ' -f2`
.GIT_HASH:=`git rev-parse --short HEAD`
.TIMESTAMP:=`date +%FT%T%z`
.LDFLAGS:=""
.BIN_NAME:=`basename "$(CURDIR)"`
export MYHOST = $(shell ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1' | awk '{print $1}')

all:
	${MAKE} getdeps
	${MAKE} clean
	govendor install -a ./...
	${MAKE} test

install:
	${MAKE} all

build:
	${MAKE} getdeps
	${MAKE} clean
	govendor build ${.LDFLAGS} ./...
	${MAKE} test

vendor:
	go get -u github.com/kardianos/govendor
	govendor sync
	govendor update

getdeps:
	go get -u github.com/kardianos/govendor
	go get -d -t ./...

docker-build:
	docker build -t ${.BIN_NAME}:${.BRANCH} -t ${.BIN_NAME}:${.VERSION} -t ${.BIN_NAME}:${.GIT_HASH} --build-arg name=${.BIN_NAME} .

docker-run:
	docker run -p 8082:8080 -e "SERVICE_NAME=service-two" -e "CONSUL_HOST=http://${MYHOST}:8500/v1/catalog/service" -it ${.BIN_NAME}:${.VERSION}

test:
	govendor test -parallel 10 -cover ./...

clean:
	govendor clean
