#! /bin/sh -e

CNTR=dredis

docker exec -it $CNTR redis-cli

# select 0 - default database
