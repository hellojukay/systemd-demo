#!/usr/bin/env bash

go build -o now main.go
sudo cp now /usr/bin/
sudo cp now.service /lib/systemd/system/now.service

# system enable now 
#  /etc/systemd/system/multi-user.target.wants/now.service → /lib/systemd/system/now.service