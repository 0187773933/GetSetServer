#!/bin/bash
APP_NAME="public-get-set-server"
sudo docker rm -f $APP_NAME || echo ""
id=$(sudo docker run -dit \
--name $APP_NAME \
--restart='always' \
--mount type=bind,source=$(pwd)/config.json,target=/home/morphs/config.json \
-v $(pwd)/SAVE_FILES/:/home/morphs/SAVE_FILES/ \
-p 5254:5254 \
$APP_NAME /home/morphs/config.json)
sudo docker logs -f $id