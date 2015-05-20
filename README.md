## Go в примерах

В этом репозитории контент, набор инструментов для сборки статического сайта и веб-сервер для [Go в примерах](http://gobyexample.ru).


### Обзор

Сайт *Go в примерах* собирается путем извлечения кода и коментариев из файлов с исходниками в папке `examples` и дальнейшим рендерингом этих данных через шаблоны в папке `templates`. Программы для сборки находятся в папке `tools`.

Процесс сборки создает папку со статическими файлами -
`public` - подходящими для обслуживания любым современным HTTP-сервером.
Мы включили в поставку легковесный Go-сервер в файле `server.go`.

### Русская версия

Русская версия была немного доработана. Теперь статику в папке `public` можно хостить прямо на Github. Подробнее о том [как настроить сборку проекта и хостинг cтатики на Github](http://gohugo.io/tutorials/github-pages-blog/).


### Сборка

Для сборки сайта:

Если вы работаете в OS X, для корректной работы сборщика сайта нужно установить Python (нужен для Pygmentize)

```console
brew install python
```

```console
$ go get github.com/russross/blackfriday
$ tools/build
$ open public/index.html
```

Для непрерывной сборки в цикле:

```console
$ tools/build-loop
```


### Локальный деплой

```bash
$ mkdir -p $GOPATH/src/github.com/pavlik
$ cd $GOPATH/src/github.com/pavlik
$ git clone git@github.com:pavlik/gobyexample.git
$ cd gobyexample
$ go get
$ foreman start
$ foreman open
```


### Деплой на платформе

Простая установка:

```bash
$ export DEPLOY=$USER
$ export APP=gobyexample-$USER
$ heroku create $APP -r $DEPLOY
$ heroku config:add -a $APP
    BUILDPACK_URL=https://github.com/mmcgrana/buildpack-go.git
    CANONICAL_HOST=$APP.herokuapp.com \
    FORCE_HTTPS=1 \
    AUTH=go:byexample
$ heroku labs:enable dot-profile-d -a $APP
$ git push $DEPLOY master
$ heroku open -a $APP
```

С добалением домена + SSL:

```bash
$ heroku domains:add $DOMAIN
$ heroku addons:add ssl -r $DEPLOY
# order ssl cert for domain
$ cat > /tmp/server.key
$ cat > /tmp/server.crt.orig
$ curl https://knowledge.rapidssl.com/library/VERISIGN/ALL_OTHER/RapidSSL%20Intermediate/RapidSSL_CA_bundle.pem > /tmp/rapidssl_bundle.pem
$ cat /tmp/server.crt.orig /tmp/rapidssl_bundle.pem > /tmp/server.crt
$ heroku certs:add /tmp/server.crt /tmp/server.key -r $DEPLOY
# add ALIAS record from domain to ssl endpoint dns
$ heroku config:add CANONICAL_HOST=$DOMAIN -r $DEPLOY
$ heroku open -r $DEPLOY
```


### Лицензия

Копирайт на изначальный проект принадлежит Mark McGranaghan и распространяется под лицензией
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

Копирайт на Go Gopher принадлежит [Renée French](http://reneefrench.blogspot.com/) и распространяется под лицензией
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).


### Переводы

Переводы с английской версии сайта Go by Example доступны на:

* [Русском](http://gobyexample.ru/) by [pavlik](https://github.com/pavlik)
* [Китайском](http://everyx.github.io/gobyexample/) by [everyx](https://github.com/everyx)
*  [Английский](http://gobyexample.com/) оригинальный вариант


### Спасибо

Спасибо [Jeremy Ashkenas](https://github.com/jashkenas)
за [Docco](http://jashkenas.github.com/docco/), вдохновившим на этот проект.
