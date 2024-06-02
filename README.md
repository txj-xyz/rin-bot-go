# RinSwaps Reactions Bot
This bots purpose is for Reactions on a review channel

Lightweight and very efficient written in GoLang


# Usage

Required variables are the following inside of the root dir of the executable.

```env
TOKEN=BotTokenString
MESSAGE_REACT_CHANNEL=ChannelID
```


# Running the Bot


Either compile or run the bot from runtime

```bash
# Run the bot in runtime
go run main.go 

# Build the bot out and run it as a job or in the background
go build -o rin-react-bot main.go 
./rin-react-bot > rin-react-bot.log 2>&1 &
```
