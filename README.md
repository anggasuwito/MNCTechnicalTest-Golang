#### PREPARATION

````
- run redis in your local
- import sql/db_technicaltest.sql 
    or 
  run for auto migrate database and then create new record in table "customer"
  created_date and updated_date can be null
````

#### RUN

````
go run main.go
````

#### API

````
- login (POST)
localhost:8085/api/customer/login
    
    example body:
    {
        "email": "angga@gmail.com",
        "password": "123"
    }

- logout (POST)
localhost:8085/api/customer/logout

- payment (POST)
localhost:8085/api/transaction/payment
    
    example body: 
    {
        "email": "admin@gmail.com",
        "bank_name": "bank syariah",
        "account_number": 9000999887,
        "amount": 1000000
    }
````