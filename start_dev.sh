#!/usr/bin/env bash

docker-compose up -d --build parking_lot
docker logs -f parking_lot