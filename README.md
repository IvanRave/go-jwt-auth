# JWT authentication without passwords

## Run

./run-docker-redis.sh
./run-app.sh

## Steps

1. Enter a email or phone number to code.html
2. Receive a verification code by email or phone, eg. 5 digits
3. Enter a email/phone + code to login.html
4. Authentication token JWT is generated and saved to a client's browser
5. A user can view secured pages


## Principles

- No passwords
- No links in emails
- Temporary codes
  - send by email or phone
  - easy remember
  - removed after first using
- JWT tokens
  - no additional storage
- Verify the phone number or email

## Stack

- Redis
- GoLang
  - jwt-go: https://github.com/dgrijalva/jwt-go
  - go-redis: https://github.com/go-redis/redis