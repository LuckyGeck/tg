# Telegram Bots
Experiments with Telegram Bot API.

## Telegram API Token

First, you need to create your new bot in Telegram by talking to [@BotFather](https://telegram.me/botfather).

Full doc: https://core.telegram.org/bots#3-how-do-i-create-a-bot

After you went through the registration, *@BotFather* sends you an HTTP API token. Copy it and pass as a command line parameter or as an environmental variable to the program below.

## Run

```shell
go run cmd/bot/main.go --token="YOUR.TOKEN"
```

or 

```shell
TGAPI_TOKEN="YOUR.TOKEN" go run cmd/bot/main.go
```

or

```shell
# run once in a terminal session or add to ~/.zshrc
export TGAPI_TOKEN="YOUR.TOKEN"

# now you can run simply
go run cmd/bot/main.go
```

Now you can talk to your telegram bot in Telegram.