#! /bin/sh -e

# download latest redis
VRSN=3.2.2
CNTR=dredis

# git pull redis:$VRSN

docker stop $CNTR || true
docker rm $CNTR || true

docker run \
       --name $CNTR \
       -p 6379:6379 \
       -d \
       redis:$VRSN

docker logs -f $CNTR
