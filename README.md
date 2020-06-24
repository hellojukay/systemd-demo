# systemd-demo
a systemd demo project , write by golang ,a web servoice return current time string.

[![Build Status](https://travis-ci.org/hellojukay/systemd-demo.svg?branch=master)](https://travis-ci.org/hellojukay/systemd-demo)


## How to install it
it will build main.go , and create now.service in `/lib/systemd/system/now.service`
```shell
sudo bash install.sh
```

## How to run it
after install now.service
```shell
systemctl start now.service
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
run with special user
```shell
vagrant@ubuntu-bionic systemd-demo (master) $ ps aux | grep now
vagrant   3604  0.0  0.1 1008360 5420 ?        Ssl  03:26   0:00 /usr/bin/now
vagrant   3866  0.0  0.0  14856  1004 pts/0    R+   03:38   0:00 grep --color=auto now
```
access by curl 
```shell
hellojukay@local systemd-demo (master) $ curl http://127.0.0.1:6000/now
2019-05-06 18:39:25
```
## Auto start at boot
it will auto start whill system started.
```shell
systemctl enable now
```
