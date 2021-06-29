

build:
	sudo docker build --no-cache . -t "civilwardebug"

run:
	sudo docker container run civilwardebug
