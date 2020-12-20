# How to link a folder with an existing Heroku app
## git:remote -a historykeeper


# local build process

create .env file
## heroku config:get db_database -s  >> .env
## heroku config:get db_host -s  >> .env
## heroku config:get db_port -s  >> .env
## heroku config:get db_database -s  >> .env
## heroku config:get db_username -s  >> .env


Procfile definition↓
## go build -o bin/hello_golang_on_heroku -v .
created bin folder
## heroku local web

goto http://localhost:5000/
