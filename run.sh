#! /bin/sh

/usr/bin/auth-server &
nginx -g "daemon off;"
