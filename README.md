# Zerocool

*My take on what used to be `z.ero.cool 1337`*

..with added fortunes angle.

Build zerocool

    docker run -it --rm -v $PWD:/go -w="/go" -e CGO_ENABLED=0 golang:1.8 go build -o bin/zerocool -a -ldflags '-s' main.go

Build the docker image

    docker build -t chanux/zerocool .

Run zerocool docker container

    docker run -it --rm -p 1337:1337  chanux/zerocool -f fortunes.txt -d 20 -p 1337  

Check it out

    nc localhost 1337

There is probably a demo at

    nc chanux.me 1337
