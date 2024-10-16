#step to use
docker build -t koriebruh/apitest .

docker compose create
docker compose start
docker compose stop

docker compose down --rmi all
