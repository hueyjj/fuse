MODULE_NAME 		= sundermodule
MODULE_BINARY_PATH 	= bin/$(MODULE_NAME)

all: fuse

fuse:
	echo "Nothing right now"
	
buildmodule:
	go build -o $(MODULE_BINARY_PATH) cmd/$(MODULE_NAME)/$(MODULE_NAME).go
