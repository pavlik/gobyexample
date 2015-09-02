# Выполнение этой программы показывает, что первая
# операция завершилась по тайм-ауту, а вторая
# выполнилась успешно
$ go run timeouts.go
timeout 1
result 2

# Using this `select` timeout pattern requires
# communicating results over channels. This is a good
# idea in general because other important Go features are
# based on channels and `select`. We'll look at two
# examples of this next: timers and tickers.
# Использование этого шаблона `select` тайм-аута требует
# обмен результатами через каналы. В общем, это хорошая
# идея, так как другие важные особенности Go основаны
# на каналах и операторе `select`. Мы ещё увидим это
# на двух примерах: таймера и счетчика тиков.
