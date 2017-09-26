# Zerocool

*My take on what used to be `z.ero.cool 1337`*

..with added fortunes angle.

Build and run 

    make build

    build bin/zerocool_linux -f fortunes.txt -d 20 -p 1337 

Run zerocool docker container (build or just pull from docker hub)

    make build-docker

    docker run -it --rm -p 1337:1337  chanux/zerocool -f fortunes.txt -d 20 -p 1337  

Check it out

    nc localhost 1337

There is probably a demo at

    nc chanux.me 1337
