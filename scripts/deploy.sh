#!/bin/bash

if [ $# -lt 1 ]; then
  echo "Use: deply.sh source_dir"
  exit 1
fi

baseProjectDir=$1
#baseProjectDir="/root/gosrc/src/main/"

binaryFile="main"
confFile="config.json"
serviceFile="fibonaccid"
monitorFile="monitorfibonaccid"

#By default use /usr/local/bin, this can be changed as an input parameter
targetDir="/usr/local/bin/"

#Default is localhost, can be changed as an input parameter
targetHost="" 

#deploy to target, now only deploy to localhost, 
#use scp to upload to remote hosts

echo "Copying files ......"
cp ${baseProjectDir}/${binaryFile} ${targetDir}fibonacci
chmod 755 ${targetDir}fibonacci
cp ${baseProjectDir}/${confFile} ${targetDir}
cp ${baseProjectDir}/scripts/${monitorFile} ${targetDir}
chmod 755  ${targetDir}${monitorFile} 
cp ${baseProjectDir}/scripts/${serviceFile} /etc/init.d
chmod 755 /etc/init.d/${serviceFile}

#add fibonaccid service
echo "Starting service ......"
chkconfig --add fibonaccid

sleep 5
#start fibonaccid service
service fibonaccid start

#add the monitor fibonaccid crontab
echo "Starting monitor job......"
echo "5 * * * * /usr/local/bin/monitorfibonaccid" > /tmp/fibonacci.cron
crontab /tmp/fibonacci.cron
service crond restart
rm -f /tmp/fibonacci.cron


