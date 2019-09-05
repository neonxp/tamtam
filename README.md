# TamTam Go

[![Sourcegraph](https://sourcegraph.com/github.com/neonxp/tamtam/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/neonxp/tamtam?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/neonxp/tamtam)
[![Go Report Card](https://goreportcard.com/badge/github.com/neonxp/tamtam?style=flat-square)](https://goreportcard.com/report/github.com/neonxp/tamtam)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/neonxp/tamtam/master/LICENSE)

Простая реализация клиента к TamTam Bot API на Go. Поддерживается получение обновление как с помощью вебхуков, так и лонгполлингом.

Поддерживаемая версия API - 0.1.8

## Документация
В общем случае, методы повторяют такие из [официальной документации](https://dev.tamtam.chat/)
Так же добавлены хелпер для создания клавиатуры (`api.Messages.NewKeyboardBuilder()`) и для загрузки вложений (`api.Uploads.UploadMedia(uploadType UploadType, filename string)`). 

Пример создания клавиатуры см. ниже в примере.
 
Остальное описано тут http://godoc.org/github.com/neonxp/tamtam/ и в примерах из директории [examples](https://github.com/neonxp/tamtam/tree/master/examples)

## Пример

* [Пример с отправкой фото](https://github.com/neonxp/tamtam/blob/master/examples/example.go)
* [Пример с longpolling](https://github.com/neonxp/tamtam/blob/master/examples/example_longpolling.go)
* [Пример с webhook](https://github.com/neonxp/tamtam/blob/master/examples/example_webhook.go)

## Автор

Александр NeonXP Кирюхин  <a.kiryukhin@mail.ru>
