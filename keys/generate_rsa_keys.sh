openssl genrsa -out auth.key 4096
openssl rsa -in auth.key -out auth.pub -pubout -outform PEM
sleep 2
openssl genrsa -out server.key 4096
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 365