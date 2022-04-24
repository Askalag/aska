module github.com/Askalag/aska/microservices/auth

go 1.17

replace github.com/Askalag/protolib v1.0.1 => ../../protolib

require (
	github.com/Askalag/protolib v1.0.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang-migrate/migrate/v4 v4.15.1
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.0
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	go.uber.org/atomic v1.6.0 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20211013075003-97ac67df715c // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211013025323-ce878158c4d4 // indirect
	google.golang.org/grpc v1.43.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
