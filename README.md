# Study Asynq

- Go: 1.18
- Asynq: v0.23.0
- Asynqmon: v0.7.1

<br/>

## Run Server
### With Docker
- For MacOS
```bash
$ runserver.sh
```

### With a local binary
```bash
# 1. run redis
$ docker run -p 6379:6379 redis 

# 2. compile
$ make build

# 3. run the server for MacOS. Or You can use
# the debug mod of VSCode to click the button 'Launch' 
$ ./bin/study-event-go-asynq
```

<br/>

## Test

### Send a task
#### With cURL
```bash
# with Docker
curl --location --request POST 'http://0.0.0.0:14569/api/v1/study-asynq/announcement/schedule' \
--header 'Content-Type: application/json' \
--data-raw '{
    "from": "The kirin",
    "message": "WAKARIMASU"
}'

# with a local binary
$ curl --location --request POST 'http://127.0.0.1:4569/api/v1/study-asynq/announcement/schedule' \
--header 'Content-Type: application/json' \
--data-raw '{
    "from": "The kirin",
    "message": "WAKARIMASU"
}'
```

#### With Postman
- Method
    - POST
- URL
    - {{url}}/api/v1/study-asynq/announcement/schedule
        - {{url}} for Docker : http://0.0.0.0:14568 or http://0.0.0.0:14569
        - {{url}} for a local binary : http://127.0.0.1:4569
- Body
    ```bash
    {
        "from": "The kirin",
        "message": "WAKARIMASU"
    }
    ```

### Result on the console
```bash
[ANNOUNCEMENT] GOT A NEW ANNOUNCEMENT FROM The kirin.
[ANNOUNCEMENT] THE MESSAGE IS ...
[ANNOUNCEMENT] "WAKARIMASU"
```

### To Check the result on the Web UI
#### With Docker
- [http://0.0.0.0:8080/monitoring](http://0.0.0.0:8080/monitoring)
#### With a local binary
- [http://127.0.0.1:8080](http://127.0.0.1:8080)
