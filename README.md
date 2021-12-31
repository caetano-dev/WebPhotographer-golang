## Web Photographer - golang version - link based branch

A Telegram bot that is able to take screenshots of webpages remotely and send it to its users.

## What is different in this branch?

This branch uses a third party service in order to provide the screenshots, therefore it is faster, lighter and simpler. It does not require a web browser in order to work.

## How to set up

In order to run the bot you will need a [Telegram](https://core.telegram.org/api) API key.

Rename the file from `example.env` to `.env` and put the key in the file.

```env
TELEGRAM_KEY=<TOKEN>
```

install packages:


```
$ go get
```


run the application:


```
$ go run main.go
```


or if you want the executable file:


```
$ go build main.go
$ ./main
```
