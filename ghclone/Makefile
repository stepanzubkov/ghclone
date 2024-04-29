PREFIX=/usr

all: ghclone

ghclone:
	go build -o ghclone main.go

install: ghclone
	install -v ghclone ${PREFIX}/bin
