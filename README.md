<p align="center"> <img src="https://miro.medium.com/max/900/1*r72eZxfLab9ZQP6R6xJ1tg.png"> </p>

# Go Gin

## Curl Requests

<b> Add Video </b>

```shell
curl --location --request POST 'http://localhost:8080/videos' \
--header 'Authorization: Basic Z29naW46Z29naW5wd2Q=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Cool",
    "description": "tokyo",
    "url": "https://s3.amazon.com/bucket/drifts/tokyo/a23kaje324n.23.mkv",
    "author": {
        "firstname": "kartik",
        "lastname": "rawat",
        "age": 22,
        "email": "k@gmail.com"
    }
}'
```

<b>Get All videos</b>

```shell
curl --location --request GET 'localhost:8080/videos'
```
