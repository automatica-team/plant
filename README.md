# 🏭
# plant - automatic bot farm, create Telegram bot boilerplate

# **Installation**:

```sh
$ go install automatica.team/plant/cmd/plant@latest
```
## **Example plant.yml file**
```yml
bot:
  token: $TOKEN # Token of your Telegram bot
  expose: # Handlers to use
    - /start
    - /t

deps:
  - import: plant/db # Importing database dependency
    dsn: $DB_URL # DSN of your DataBase

mods:
  - import: plant/core # Importing core module
  - import: x/tracker # Importing private (x) module tracker
```
## **Usage**:

```plant [OPTIONS] COMMAND [ARGS]```
**Available commands**:
```
  build       Builds a Docker image for the bot
  help        Help about any command
  run         Create and run a new bot from a config
  version     Print version and quit
```
**Credits**:
