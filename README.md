# authentication-service
This is a repository for an assessment.

#### Install postgres and setup db from bash console while in repository root directory:
```
createdb auth_service_db
psql auth_service_db < auth_dump.sql
```

### SERVICE PORT: `9000`

### Dependency:
```
otp_service at PORT 8000
test_key.json file for google pub sub authentication, it will be provided by me in same mail where assignment was given, will deactivate it after the feedback
```

### Generating proto buffers from root directory of repository(Already checked in so no need to generate):
`protoc --go_out=plugins=grpc:proto_buffers --go_opt=paths=source_relative ./protos/*.proto`

### Running the server:
`go run ./server/server.go`

### Running the signup_client(gRPC ACTION):
`go run ./clients/signup_client/signup_client.go`

### Running the verify_client(gRPC ACTION use it after signup client and login client to verify OTP):
`go run ./clients/verify_client/verify_client.go`

### Running the login_client(gRPC ACTION):
`go run ./clients/login_client/login_client.go`

### Running the logout_client(gRPC ACTION):
`go run ./clients/logout_client/logout_client.go`

### Running the get_user_client(gRPC ACTION):
`go run ./clients/get_user_client/get_user_client.go`

### NOTE:
`Please fill the user form data like name, phone_number, otp, etc accordingly in all the clients`
