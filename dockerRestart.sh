#!/bin/bash
APP_NAME="public-get-set-server"
id=$(sudo docker restart $APP_NAME)
sudo docker logs -f $id