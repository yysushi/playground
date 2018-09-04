```
docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_USER=user -e MYSQL_PASSWORD=password -e MYSQL_DATABASE=database -v "$(pwd)/init:/docker-entrypoint-initdb.d" --rm -d mysql:5.5
```
