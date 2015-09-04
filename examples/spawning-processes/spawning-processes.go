// Иногда, вашим Go программ необходимо вызвать другие,
// не Go, процессы. Например, подсветка синтаксиса на этом
// сайте [реализовано](https://github.com/pavlik/gobyexample/blob/master/tools/generate.go)
// при помощи порождения [`pygmentize`](http://pygments.org/)
// процесса из Go программы. Давайте рассмотрим несколько
// примеров порождения процессов из Go.

package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {

    // Мы начнем с простой команды у которой нет входных
    // аргументов или параметров и она только выводит что-то в
    // стандартный поток вывода (stdout). Команда `exec.Command`
    // является объектом, который предоставляет доступ к этому внешнему процессу.
    dateCmd := exec.Command("date")

    // `.Output` - команда которая ждет выполнения и
    // завершения процесса, сохраняя его вывод.
    // Если не произошло ошибки, `dateOut` будет заключать в себе
    // байты с информацией о дате.
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))

    // В этом примере мы заглянем чуточку глубже, в случай
    // когда мы направляем данные во внешний процесс
    // в его поток ввода (stdin) и считываем результаты
    // из потока вывода этого процесса.
    grepCmd := exec.Command("grep", "hello")

    // Здесь мы непосредственно берем потоки (pipes) ввода/вывода,
    // стартуем процесс, записываем в него некоторые входные
    // данные, читаем результирующий выходные данные и
    // ждем пока процесс завершится.
    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()

    // Мы опустили проверку на ошибки в примерах выше, но
    // вы можете использовать для этого шаблон `if err != nil`.
    // Так же мы сохраняем данные только из потока вывода (StdoutPipe),
    // но вы можете сохранять из потока ошибок (`StderrPipe`)
    // в точности также.
    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

    // При порождении команд нам необходимо предоставить
    // четко определенную команду и массив ее аргументов,
    // вместо того чтобы задать всё в одной строке. Если
    // вы хотите породить команду в одной строке, то вы
    // можете использовать `bash` `-c`:
    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}
