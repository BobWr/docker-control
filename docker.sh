#!/bin/bash


#if deal equals 0, run new container
#else if deal equals 1, stop and remove container
DEAL=$1

PORT=$2

if [ $DEAL -eq 0 ]
then
  echo "[log]run container zukalp$PORT"
  docker run -p $PORT:9999 --cap-add=SYS_ADMIN --name zukalp$PORT zukrec_alpine_29
else
  echo "[log]rm container zukalp$PORT"
  docker stop zukalp$PORT
  docker rm zukalp$PORT
fi