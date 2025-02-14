#! /bin/bash
docker build -t forum-app .
docker run -d -p 8080:8080 forum-app
docker system prune -f