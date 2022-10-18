#!/bin/bash

echo "waiting db to accept connections..."
deployment/wait-for-it.sh -t 0 db:5432
sleep 5

echo "creating database and doing migrations..."
(buffalo pop create -a; buffalo pop migrate) 2> /dev/null

echo "starting application..."
/bin/app
