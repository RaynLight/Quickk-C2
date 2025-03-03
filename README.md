# Quickk C2
 Repo for the 2025 deans cup at Dakota State 

### Server Setup
```
git clone https://github.com/RaynLight/Quickk-C2.git
go get github.com/google/uuid
go get github.com/quic-go/quic-go/http3

openssl genpkey -algorithm RSA -out server.key
openssl req -new -key server.key -out server.csr
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt

go run .
```
