# Snowflake Private Key Auth in Go

Securely connect to [Snowflake](https://www.snowflake.com/) using **JWT-based private key authentication** with Go. This project demonstrates a production-grade structure using:

- Private Key Auth with Snowflake
- Clean architecture: Handler → Service → Repository
- Gin for REST APIs
- Environment-based configuration
- Secure key handling (PEM + Base64 + passphrase)

---

## Project Structure
```
.
├── main.go # Application entry point
├── internal/
│ ├── factories/ # Snowflake connection factory
│ ├── repositories/ # Database layer
│ ├── services/ # Business logic
├── server/handlers/ # HTTP handlers
├── pkg/utils/ # Logger and env validation
```

## Add your .env file
To encode your PEM key:
```
cat rsa_key.p8 | base64
```
Create a .env file in the root directory with the following variables:
```
SNOWFLAKE_ACCOUNT=your_account_id
SNOWFLAKE_USER=your_user
SNOWFLAKE_WAREHOUSE=your_warehouse
SNOWFLAKE_DATABASE=your_database
SNOWFLAKE_SCHEMA=your_schema
SNOWFLAKE_PRIVATE_KEY=BASE64_ENCODED_PEM_STRING
SNOWFLAKE_PASSPHRASE=your_key_passphrase
```
## Running the Server
```
go run main.go
```

## How Authentication Works
The app connects to Snowflake using JWT-based auth.

The RSA private key is Base64 encoded, then decrypted using a passphrase at runtime.

Go’s x509 and pem libraries are used to parse and decrypt the key.

The gosnowflake driver handles the connection securely.

## 🤝 Contributing
Feel free to fork the repo, raise issues, or submit pull requests!

## Blog Post
[Connecting to Snowflake Using Private Key Auth in Go](http://example.com)