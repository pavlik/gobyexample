# Чтобы поэкспериментировать с аргументами 
# командной строки, лучше всего сначала собрать
# бинарный файл командой `go build`.
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c

# Далее мы рассмотрим более продвинутую работу с командной
# строкой через флаги.
