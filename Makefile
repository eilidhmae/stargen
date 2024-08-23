build:
	go fmt
	go fix
	go vet
	go build -o sg

clean:
	rm -f sg
