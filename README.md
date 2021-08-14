# Go Gin

## Curl Requests

<b> Add Video </b>

```
curl --location --request POST 'http://localhost:8080/videos' \
--header 'Authorization: Basic Z29naW46Z29naW5wd2Q=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "drifts",
    "description": "tokyo",
    "url": "https://s3.amazon.com/bucket/drifts/tokyo/a23kaje324n.23.mkv"
}'
```

<b>Get All videos</b>

```
curl --location --request GET 'localhost:8080/videos'
```
