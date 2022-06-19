bin/ely:
	go fmt
	go build -o bin/ely

clean:
	rm -Rf bin/ely 

build: clean bin/ely

