FILE_NAME=dgncrwlr

build:
	go build -o ${FILE_NAME}

install:
	sudo cp ${FILE_NAME} /usr/bin

remove:
	sudo rm /usr/bin/${FILE_NAME}