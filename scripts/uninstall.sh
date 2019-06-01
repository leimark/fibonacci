#!/bin/bash

service fibonaccid stop
sleep 2
chkconfig --del fibonaccid

rm -fr /var/log/gin-access.log
rm -fr /var/log/fibonacci.log

rm -fr /usr/local/bin/fibonacci
rm -fr /usr/local/bin/monitorfibonaccid
rm -fr /usr/local/bin/config.json
