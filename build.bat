docker build --tag todoapp-server:latest -f Dockerfile .
docker tag todoapp-server:latest halitemen123/todoapp-server:latest
docker push halitemen123/todoapp-server:latest
pause