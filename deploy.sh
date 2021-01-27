env GOOS=linux go build -o bot_support
ssh root@157.230.97.187 "systemctl stop bot-support.service"
scp bot_support root@157.230.97.187:~/monitors/bot_support
ssh root@157.230.97.187 "systemctl start bot-support.service"
ssh root@157.230.97.187  "systemctl status -l bot-support.service"
rm bot_support