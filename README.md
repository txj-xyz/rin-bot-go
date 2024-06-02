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


# Service files
Running the bot as a service can be done with `systemd` here are the instructions below

Install the service for systemd under `/etc/systemd/system/rin-react-bot.service`
```bash
[Unit]
Description=Rin React Bot Service
After=network.target

[Service]
ExecStart=/srv/rin-bot-go/rin-react-bot
WorkingDirectory=/srv/ubuntu/rin-bot-go
Restart=always
User=ubuntu
StandardOutput=file:/var/log/rin-bot-go/rin-react-bot.log
StandardError=file:/var/log/rin-bot-go/rin-react-bot.log

[Install]
WantedBy=multi-user.target
```

Reload systemd after installation
`sudo systemctl daemon-reload`

Start up the Service
`sudo systemctl start rin-react-bot.service`

Start the service on boot
`sudo systemctl enable --now rin-react-bot.service`
