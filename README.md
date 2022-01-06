# SETUP

1. Change your local postgresql configs in `.env` file
2. Connect with posgresql and Create `database` and `users` table manually.

```
i) CREATE DATABASE gocrud;
ii) CREATE TABLE users (userid SERIAL PRIMARY KEY,name TEXT,age INT,location TEXT);
```

3. Run the server - `go run main.go`

## Test APIs with Postman

Use [this](https://github.com/parthivpatel21/go-crud/blob/master/documents/goCRUD.postman_collection.json) postman collection