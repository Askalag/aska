Under development

Go microservices+

.01 protolib/ $ buf generate

## Auth
### Docker 
docker build --tag auth-app .

### Генерация rsa private/public
openssl genrsa -out privat_rsa.pem 2048
openssl rsa -in privat_rsa.pem -out public_rsa.pem -pubout -outform PEM
