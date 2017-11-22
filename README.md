# Genderize Server

Super simple API for first name to gender conversion.
Currently 16.201 Turkey, 31.015 USA first name included (see data folder).

## Install
```bash
go get github.com/ismailbaskin/genderize
```

## Run
```bash
PORT=9090 genderize
```
Note: default port is 8080

## Docker
```bash
docker run -p 8080:8080 -d ismailbaskin/genderize
```

## HTTP Usage

```bash
curl 'http://127.0.0.1:9090/Mustafa'
```

Result

```json

[
    {
        "name": "Mustafa",
        "gender": "male",
        "accuracy": 100
    }
]
```

#### Multi name usage

```bash
curl 'http://127.0.0.1:9090/Halil%20İbrahim'
```

Result
```json
[
    {
        "name": "Halil",
        "gender": "male",
        "accuracy": 100
    },
    {
        "name": "İbrahim",
        "gender": "male",
        "accuracy": 100
    }
]
```

#### JSONP Support
```bash
curl 'http://127.0.0.1:9090/Halil%20İbrahim?callback=mycallback'
```

Result
```javascript
mycallback([{"name":"halil","gender":"male","accuracy":100},{"name":"İbrahim","gender":"male","accuracy":100}])
```
