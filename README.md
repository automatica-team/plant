# 🏭
**Plant** is a declarative Telegram Bot builder built using [Telebot](https://github.com/tucnak/telebot).

## Quickstart

```bash
$ go install automatica.team/plant/cmd/plant@latest
$ plant init example
$ plant run
```

## Config
```yml
bot:
  expose:
    - /start
    - /t

deps:
  - import: plant/db
    dsn: $DB_URL

mods:
  - import: plant/core
  - import: x/private
```
