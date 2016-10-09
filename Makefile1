myapp: myapp.go
	docker run -v $(PWD):/src -w /src golang \
	  go build -tags netgo -installsuffix netgo -o myapp myapp.go
