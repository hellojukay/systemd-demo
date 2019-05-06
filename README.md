# systemd-demo
a systemd demo project , write by golang ,a web servoice return current time string.
## how to install it
it will build main.go , and create now.service in `/lib/systemd/system/now.service`
```shell
sudo bash install.sh
```

## how run it
after install now.service
```shell
system start now.service
```
```shell
hellojukay@local systemd-demo (master) $ systemctl status now
● now.service - A web service get now time string
   Loaded: loaded (/lib/systemd/system/now.service; enabled; vendor preset: enabled)
   Active: active (running) since Mon 2019-05-06 17:52:35 CST; 5min ago
 Main PID: 3052 (now)
    Tasks: 6 (limit: 4915)
   Memory: 5.4M
   CGroup: /system.slice/now.service
           └─3052 /usr/bin/now

Warning: Journal has been rotated since unit was started. Log output is incomplete or unavailable.
```
## auto start at boot
it will auto start whill system started.
```shell
system enable now
```