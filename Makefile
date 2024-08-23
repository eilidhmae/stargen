build:
	go fmt
	go fix
	go vet
	go build -o stargen

clean:
	rm -f stargen
