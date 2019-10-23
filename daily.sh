#!/data/data/com.termux/files/usr/bin/env bash

cd ~/daily-warm
cp .env.daily .env
echo $(date "+%n%Y-%m-%d %H:%M:%S")
./dwm.out 
rm .env





