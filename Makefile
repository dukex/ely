output/bin/ely:
	go build -o output/bin/ely ./cmd/ely

clean:
	rm -Rf output/bin/ely 

build: clean output/bin/ely

build_and_run: build 
	DATABASE_URL=postgresql://localhost:5432/ely?sslmode=disable ./output/bin/ely db setup
	DATABASE_URL=postgresql://localhost:5432/ely?sslmode=disable ./output/bin/ely server -p 3000

