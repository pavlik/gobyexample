# Чтобы проверить наш линейный фильтр, создайте
# файл с парой строк в нижнем регистре.
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# Затем, используйте программу, чтобы получить
# капитализированные строки.
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
