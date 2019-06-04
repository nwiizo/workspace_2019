#!/bin/bash
cp test.sh /tmp/test.sh
chmod 770 /tmp/test.sh
echo "* * * * * root /tmp/test.sh"> /etc/cron.d/rubykaigi_day01
ls -l /tmp/test.sh
cat /etc/cron.d/rubykaigi_day01
