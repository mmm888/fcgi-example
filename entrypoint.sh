#!/bin/sh

nginx
/fcgiapi &

tail -f /dev/null
