## How to link a folder with an existing Heroku app
```
git:remote -a historykeeper
```
----------------------------

## build process
```
govendor init
govendor fetch +out
```

1. Procfile 
```
web: bin/history_keeper
```
2.vendor/vendor.json
```
 "rootPath": "history_keeper"
```
3. deploy app
```
git push -u origin main
```
4. open app
```
heroku open 
```
----------------------------

### local build process

1. create .env file
```
touch .env
heroku config:get db_database -s  >> .env
heroku config:get db_host -s  >> .env
heroku config:get db_port -s  >> .env
heroku config:get db_database -s  >> .env
heroku config:get db_username -s  >> .env
```
2. created bin folder Procfile definition↓
```
go build -o bin/history_keeper -v .
```
3. deploy app
```
heroku local web
```
4. open app
```
open http://localhost:5000/
```
