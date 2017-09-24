from scratch

add bin/zerocool /zerocool
add fortunes.txt /fortunes.txt

expose 1337
entrypoint ["/zerocool"]
cmd ["-f", "/fortunes.txt"]
