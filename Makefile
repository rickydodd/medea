.PHONY: install-go update-go

install-go:
	make update-go
	echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc

update-go:
	curl -OL https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
	sudo rm -rf /usr/local/go
	sudo tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz
	rm go1.22.2.linux-amd64.tar.gz

.PHONY: tidy

tidy:
	go mod tidy
	go fmt ./...

.PHONY: build

build:
	make tidy
	go build -o /tmp/medea ./main.go
