# Чтобы поэкспериментировать с флагами командной
# строки, лучше всего сначала скомпилировать код,
# а затем запускать бинарный файл напрямую.
$ go build command-line-flags.go

# Попробуйте запустить собранную программу, сначала
# указав значения для всех флагов.
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# Заметьте, если вы пропустите какие-то флаги,
# они будут равны значениям по-умолчанию.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# Оставшиеся элементы командной строки будут выведены
# в конце.
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# Обратите внимание, что пакет `flag` требует,
# чтобы все флаги были указаны до оставшихся элементов
# командной строки. В противном случае, флаги будут
# обработаны как оставшиеся аргументы.
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
trailing: [a1 a2 a3 -numb=7]

# Используйте флаги `-h` или `--help`, чтобы вывести
# на экран автоматически сгенерированное текст с
# описанием флагов.
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Если вы укажете флаг, который не был указан в программе,
# то `flag` выдаст сообщение об ошибке и текст со всеми
# возможными флагами.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...

# Далее мы познакомимся с переменными окружения, еще одним
# способом параметризации программ.
