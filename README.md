# struktur-non-marketing
rest-api for service struktur non marketing

## Readme

### clone repo 
``` git
git clone https://github.com/desalisra/struktur-non-marketing.git
```

### install dependency
``` go
go mod tidy
```

### running program 
``` go
go run cmd/http/main.go
```


## List Endpoint

base_url : localhost:8080/struktur-non-mkt/

# get data master
## /master

### /departments (GET)
### /department/{id} (GET)
(example params)
``` js
/department/227
```


### /position/{id} (GET)
(example params)
``` js
/position/281
```
### /cities (GET)
### /city (GET)
(example query)
``` js
type = ["id", "name"]
value = "jakarta"
```
### /branch/{id} (GET)
(example params)
``` js
/branch/2
```


### /jabatan-iklan (GET)
(example query)
``` json
pt_id = 1
jab_id = 281
dpt_id = 227
```


# get data struktur
## /struktur
### /groupteri (GET)
(Example query)
``` json
pt_id = 1
dpt_id = 227
```
