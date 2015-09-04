# Порожденная команда возвращает в точности такой же
# результат, если бы мы запустили ее непосредственно
# в командной строке.
$ go run spawning-processes.go 
> date
Wed Oct 10 09:53:11 PDT 2012

> grep hello
hello grep

> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go
