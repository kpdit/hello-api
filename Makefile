GO_VERSION := 1.16.6  #

setup: #
	install-go
	init-go

install-go: #
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

init-go: #
	echo 'export PATH=$$PATH:/usr/local/go/bin' >> $${HOME}/zshrc
	echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.zshrc

build: #
	go build -o api cmd/main.go 

