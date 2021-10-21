.PHONY: build
build:
	docker build -t httpserver:v1 .
    
.PHONY: clean
clean:
	-docker stop httpserver
	-docker rm httpserver
	-docker rmi httpserver:v1
	
run: build
	docker run -d --name httpserver  -p 80:8080 httpserver:v1 
