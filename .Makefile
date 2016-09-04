myapp: myapp.go
	docker run -v $(PWD):/src -w /src golang \
	  go build -o myapp myapp.go
