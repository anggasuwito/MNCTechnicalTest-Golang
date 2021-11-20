#### PREPARATION

````
- edit .env file
- run redis in your local
- create database with name db_technicaltest (initial customer will be added automatically)
````

#### RUN

````
go run main.go
````

#### TEST
````
go test ./...
````

#### CUSTOMER ACCOUNT
````
- customer 1
    email : customer1@gmail.com
    password : 123

- customer 2
    email : customer2@gmail.com
    password : 123
````

#### API

````
- login (POST)
localhost:8085/api/customer/login
    
    example body:
    {
        "email": "customer1@gmail.com",
        "password": "123"
    }

- logout (POST)
localhost:8085/api/customer/logout

- payment (POST)
localhost:8085/api/transaction/payment
    
    example body: 
    {
        "email": "customer1@gmail.com",
        "bank_name": "bank syariah",
        "account_number": 9000999887,
        "amount": 1000000
    }
````