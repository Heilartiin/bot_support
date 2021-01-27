env GOOS=linux go build -o bot_support
ssh root@64.225.107.99 "systemctl stop bot_support.service"
scp bot_support root@64.225.107.99:~/bot_support/bot_support
ssh root@64.225.107.99 "systemctl start bot_support.service"
ssh root@64.225.107.99  "systemctl status -l bot_support.service"
rm bot_support