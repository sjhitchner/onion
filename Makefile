PROJECT_NAME=onion-tutorial
PROJECT_PATH=github.com/sjhitchner/onion-tutorial
ONION_ID=92D7


BRANCH=`git rev-parse --abbrev-ref HEAD`
SHA=`git rev-parse HEAD`
SHORT_SHA=`git rev-parse --short=7 HEAD`
MEDIUM_SHA=`git rev-parse --short=20 HEAD`
GOPATH=$(abspath $(subst $(PROJECT_PATH),,$(CURDIR)))

SSH_KEY=$(HOME)/.ssh/id_rsa.pub

# omega2-ctrl gpiomux set i2c gpio  # i2c pins for general io
# omega2-ctrl gpiomux set pwm0 pwm  # enable pwm on 18
# omega2-ctrl gpiomux set pwm1 pwm  # enable pwm on 19
# insmod w1-gpio-custom bus0=0,5,0  # enable one wire on 5
#

ONION_ENV=
GOOS=linux
GOARCH=mipsle 

LDFLAGS=-ldflags="-s -w"
BIN_DIR=bin

build: build-led

build-led:
	GOOS=linux GOARCH=mipsle go build $(LDFLAGS) -o $(BIN_DIR)/led led/led.go
	rsync -P -a $(BIN_DIR)/led root@omega-$(ONION_ID).local:/root

# call onion-build
onion-build:
	GOOS=linux GOARCH=mipsle go build -o $(BIN_DIR)/$(1) $(1)/...

onion-setup:
	git init
	go mod init $(PROJECT_PATH)
	echo "password: onioneer"
	scp $(SSH_KEY) root@omega-$(ONION_ID).local:/etc/dropbear/authorized_keys

ssh:
	ssh root@omega-$(ONION_ID).local

.PHONY: all build test clean run
