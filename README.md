- instalar docker
- en la consola poner el siguiente comando:
    docker run --name=stock-controller -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=true -d mysql:latest
- crear db
- crear archivo .env en la carpeta client con la siguiente estructura:
```
DB_USER=tu_user
DB_HOST=tu_host
DB_PORT=tu_port
DB_NAME=nombre_tu_db
```