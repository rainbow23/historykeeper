## How to link a folder with an existing Heroku app
```
git:remote -a historykeeper
```
----------------------------

## build process
1. exec below command
```
govendor init
govendor fetch +out
```
2. add line in Procfile
```
web: bin/historyKeeper
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
go build -o bin/historyKeeper -v .
```
3. deploy app
```
heroku local web
```
4. open app
```
open http://localhost:5000/
```

----------------------------

### table
```
CREATE TABLE user_info( \
    id INT NOT NULL AUTO_INCREMENT,\
    username   VARCHAR(64) NOT NULL UNIQUE,\
    password   VARCHAR(64) NOT NULL UNIQUE,\
    created_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,\
    updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,\
    index(id)\
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ;
```

```
create table shell_history2(\
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,\
    username    VARCHAR(64) NOT NULL,\
    command     VARCHAR(256) NOT NULL,\
    uuid        VARCHAR(64),\
    date        DATETIME     NOT NULL,\
    CONSTRAINT fk_username\
        foreign key (username)\
        references user_info(username)\
        ON DELETE RESTRICT ON UPDATE RESTRICT\
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ;
```
