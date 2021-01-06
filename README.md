# OmnIS Server

[![License](https://img.shields.io/badge/license-Apache%20license%202.0-blue.svg)](https://github.com/omnis-org/omnis-server/blob/main/LICENSE)

OmnIS Server is part of the OmnIS project. It allows the processing of client data and manages user authentication.

![omnis_logo](./omnis_logo.png)



## How to build ?


```bash
cd build
make build
# That will generate a binary omnis-server
```

## Generate keys

Omnis server needs key to communicate in HTTPS and to manage authentication. They can be generated as follows :

```
cd keys
./generate_rsa_keys.sh
```

## Create Database

Omnis rest api needs to store informations in a mariadb database.

You can create table and procedures with the following commands :

```bash
sudo apt install mariadb-server
sudo mysql_secure_installation

cd sql/

# If you haven't define root password :
sudo mysql < create_db.sql
sudo mysql OMNIS < create_procedure.sql

# Else :
mysql -u root -p < create_db.sql
mysql -u root -p OMNIS  < create_procedure.sql
```

> ⚠️ You should change the password of the default create users in file : sql/create_db.sql


## Create configuration file

You have examples of configuration file in build/testdata/example.json :

```
{
    "server" : {
        "ip" : "0.0.0.0",                           # The listening IP address of the server
        "port" : 4320,                              # The port of the omnis server service
        "omnis_api" : "/api/omnis",                 # The path of omnis api
        "admin_api" : "/api/admin"                  # The path of admin api
    },
    "admin":{
        "expiration_token_time" : 10,               # Token validity time (minute)
        "auth_key_file" : "../keys/auth.key",       # RSA private key for authentification
        "auth_pub_file" : "../keys/auth.pub"        # RSA public key for authentification
    },
    "omnis_db" : {
        "name" : "OMNIS",                   # Name of database that store clients data
        "username" : "omnis",               # username of user that can access database (OMNIS)
        "password" : "PASSWORD",            # password of user that can access database (OMNIS)
        "host" : "127.0.0.1",               # The IP address of the database (OMNIS)
        "port" : 3306                       # The port of the database service (OMNIS)
    },
    "admin_db" : {
        "name" : "OMNIS_ADMIN",             # Name of database that store users data
        "username" : "omnis",               # username of user that can access database (OMNIS_ADMIN)
        "password" : "PASSWORD",            # password of user that can access database (OMNIS_ADMIN)
        "host" : "127.0.0.1",               # The IP address of the database (OMNIS_ADMIN)
        "port" : 3306                       # The port of the database service (OMNIS_ADMIN)
    },
    "tls": {
        "activated" : true,                         # Activate TLS ?
        "server_key_file": "../keys/server.key",    # RSA private key for TLS
        "server_crt_file" : "../keys/server.crt"    # RSA cert file for TLS
    }
}

```

## How to launch ?

Lauch the server with the created config file :

```bash
./omnis-server testdata/example.json
```


## Licensing

OmnIS Client is licensed under the Apache License, Version 2.0. See [LICENSE](https://github.com/omnis-org/omnis-server/blob/main/LICENSE) for the full license text.
