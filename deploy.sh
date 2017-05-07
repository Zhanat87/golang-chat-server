#!/usr/bin/env bash

rm golang-chat-server
git add . && git commit -m 'deploy' && git push origin master
# stop & remove all docker containers
docker stop $(docker ps -a -q)
# remove image
docker rmi $(docker images --filter=reference='zhanat87/golang-chat-server') -f
# build, push and pull new image
docker build -t zhanat87/golang-chat-server .
docker push zhanat87/golang-chat-server
docker pull zhanat87/golang-chat-server