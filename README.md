# OmnIS Server

[![License](https://img.shields.io/badge/license-Apache%20license%202.0-blue.svg)](https://github.com/omnis-org/omnis-server/blob/main/LICENSE)

OmnIS Server is part of the OmnIS project. It allows the processing of client data and manages user authentication.

![omnis_logo](./omnis_logo.png)



## How to build ?


```bash
cd build
./build_server.sh
# That will generate a binary omnis-server
```

## Generate keys

Omnis server needs key to communicate in HTTPS and to manage authentication. They can be generated as follows :

```
cd keys
./generate_rsa_keys.sh
```


## Create configuration file

You have examples of configuration file in build/testdata/example.json :

```
{
    "server" : {
        "ip" : "0.0.0.0",                           # The listening IP address of the server
        "port" : 4320                               # The port of the omnis server service
    },
    "worker" : {
        "wait_work_time" : 10                       # Time during worker waits if it has no work to do (second)
    },
    "admin":{
        "expiration_token_time" : 10,               # Token validity time (minute)
        "auth_key_file" : "../keys/auth.key",       # RSA private key for authentification
        "auth_pub_file" : "../keys/auth.pub"        # RSA public key for authentification
    },
    "rest_api" : {
        "ip" : "127.0.0.1",                         # The IP address of the omnis rest api service
        "port" : 4321,                              # The port of the omnis rest api service
        "admin_path" : "/admin",                    # Path of the administration part
        "omnis_path" : "/api",                      # Path of the omnis data path
        "tls": true,                                # Is TLS activated ?
        "insecure_skip_verify": false               # Check if certificate is valid
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
