# Запущенная программа показывает, что выполняются 9
# задач разными обработчиками. Эта программа тратит только
# примерно 3 секунды несмотря на общую работу в 9 секунд.
# Так как 3 обработчика работают одновременно.
$ time go run worker-pools.go 
worker 1 processing job 1
worker 2 processing job 2
worker 3 processing job 3
worker 1 processing job 4
worker 2 processing job 5
worker 3 processing job 6
worker 1 processing job 7
worker 2 processing job 8
worker 3 processing job 9

real	0m3.149s
