# TechnicalTask

## Quick start

To develop/run the server locally, there is a development environment defined in `docker-compose-yaml` that can be used. This will create 2 containers:
- A development environment (defined in the Dockerfile under the `dev` target)
- A Postgres service to use

**To spin up the dev env**
```
make dev-up
```

**To enter the dev env**
```
make dev-shell
```

**To teardown the dev env**
```
make dev-down
```

### Start the server

**From within the dev env:**

To start the server, first build the project using the Make target...
```
make build
```
This will build the project in the `dist` directory within the container. The server can then be started using...
```
./dist/transactionServer
```

The server will then be running and listening for request on `localhost:3000`.

### Sending requests

The compose file used to create the dev env environment exposes the containers `3000` port to the host, so requests can be made either within the dev container or on the host machine.

For example getting an account via `curl` on the host machine...
```bash
curl --request GET \
  --url http://localhost:3000/accounts/018fdf82-1286-7b07-9c4b-073403bfddf4 \
  --header 'Accept: application/json, application/problem+json' | jq

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   123  100   123    0     0   9420      0 --:--:-- --:--:-- --:--:--  9461
{
  "account_id": "018fdf82-1286-7b07-9c4b-073403bfddf4",
  "document_number": "12345",
  "created_at": "2024-06-03T19:09:54.694723Z"
}
```

### API documentation

When the server is running (either locally or in the dev container), the API documentation can be viewed on [http://localhost:3000/docs](http://localhost:3000/docs).

Requests can also be made via this documentation site.

### Inside the development container

The following Make targets are also available:
- `make unit` - Runs the unit tests
- `make fmt` - format the code with `gofmt`
- `make dev-db-connect` - Connect to the database via `psql`
- `make mocks` - Regenerate the mocks using `mockery`

### Configuration

The configuration settings for running locally are stored in [bench-config.toml](bench-config.toml).

The app gets its configuration settings from the a toml file that should be specified using the environment variable `CONFIG_FILE`. For local execution, this is set to `bench-config.toml` within the [docker compose file](docker-compose.yaml).

## My submission

From the instuctions given, I made the following assumptions when creating the API:
- There is a 1:1 mapping between accounts/customers and the document number, meaning that an account cannot be created with the same document number as another account.
- When creating an account, we should also store its creation timestamp
- When requesting to create a transaction, all amounts are entered as positive integers in the lowest denomination. They are then made positive or negative based on the operation type before being stored in the database.

I decided to implement this project following clean architecture to improve the future support and flexibility of the app/API. 

I have added some unit tests, although I didn't have enough time to write all the tests I would have liked, hopefully what is present gives an idea of my intentions. These test make use of [mockery](https://github.com/vektra/mockery) for generating based on the interfaces within the code (e.g. for mocking usecase or database operations).

### Examples

**Creating an account**
```sh
curl --request POST \
  --url http://localhost:3000/accounts \
  --header 'Accept: application/json, application/problem+json' \
  --header 'Content-Type: application/json' \
  --data '{
  "document_number": "999"
}' | jq


  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   154  100   124  100    30  12493   3022 --:--:-- --:--:-- --:--:-- 17111
{
  "account_id": "018fdf9f-fb46-7891-b644-fab15dd91e23",
  "document_number": "999",
  "created_at": "2024-06-03T19:42:34.822569001Z"
}
```

**Getting an account**
```sh
curl --request GET \
  --url http://localhost:3000/accounts/018fdf9f-fb46-7891-b644-fab15dd91e23 \
  --header 'Accept: application/json, application/problem+json' | jq


  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   121  100   121    0     0  12676      0 --:--:-- --:--:-- --:--:-- 13444
{
  "account_id": "018fdf9f-fb46-7891-b644-fab15dd91e23",
  "document_number": "999",
  "created_at": "2024-06-03T19:42:34.822569Z"
}

```

**Creating 2 transactions**
```sh
curl --request POST \
  --url http://localhost:3000/transactions \
  --header 'Accept: application/json, application/problem+json' \
  --header 'Content-Type: application/json' \
  --data '{
  "account_id": "018fdf9f-fb46-7891-b644-fab15dd91e23",
  "amount": 150,
  "operation_type_id": 2
}' | jq


  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   290  100   189  100   101   6981   3730 --:--:-- --:--:-- --:--:-- 10740
{
  "transaction_id": "018fdfa1-c78b-7cba-a0fe-7b586d476edd",
  "account_id": "018fdf9f-fb46-7891-b644-fab15dd91e23",
  "operation_type": 2,
  "amount": -150,
  "event_date": "2024-06-03T19:44:32.651845417Z"
}

curl --request POST \
  --url http://localhost:3000/transactions \
  --header 'Accept: application/json, application/problem+json' \
  --header 'Content-Type: application/json' \
  --data '{
  "account_id": "018fdf9f-fb46-7891-b644-fab15dd91e23",
  "amount": 100,
  "operation_type_id": 4
}' | jq


  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   289  100   188  100   101  31639  16997 --:--:-- --:--:-- --:--:-- 57800
{
  "transaction_id": "018fdfa2-8d38-748a-b577-91cecc370e7b",
  "account_id": "018fdf9f-fb46-7891-b644-fab15dd91e23",
  "operation_type": 4,
  "amount": 100,
  "event_date": "2024-06-03T19:45:23.256302927Z"
}
```

**Viewing the created transaction in the database**
```
make dev-db-connect
transactions-db=# select * from transactions where account_id='018fdf9f-fb46-7891-b644-fab15dd91e23';
                  id                  |              account_id              |     operation_id     | amount |         event_time
--------------------------------------+--------------------------------------+----------------------+--------+----------------------------
 018fdfa1-c78b-7cba-a0fe-7b586d476edd | 018fdf9f-fb46-7891-b644-fab15dd91e23 | INSTALLMENT PURCHASE |   -150 | 2024-06-03 19:44:32.651845
 018fdfa2-8d38-748a-b577-91cecc370e7b | 018fdf9f-fb46-7891-b644-fab15dd91e23 | PAYMENT              |    100 | 2024-06-03 19:45:23.256302
(2 rows)
```
