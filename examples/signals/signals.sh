# Когда мы запустим эту программу она будет ожидать
# сигнал. При нажатии `ctrl-C`, мы посылаем `SIGINT`.
# При его обработке  программа выведет `interrupt`
# и завершится.
$ go run signals.go
awaiting signal
^C
interrupt
exiting
