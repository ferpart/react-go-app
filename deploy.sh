#!/bin/sh
sudo docker-compose down
sudo git pull origin dev
sudo docker-compose up --build
