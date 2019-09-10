# Documentation

## Setup

docker-compose build
docker-compose up

## API Routes

### POST "/create-account"
Send in the body the following:
```
name string
balance_amount int64
```
Create an account with a given balance amount
### POST "/credit-account"
Send in the body the following:
```
accountID string
amount int64
```
Credit the given account with the given amount
### POST "/debit-account"
Send in the body the following:
```
accountID string
amount int64
```
Debit the given account with the given amount
### POST "/make-transaction"
Send in the body the following:
```
senderID string
receiverID string
amount int64
description string
note string
```
Execute a transaction of the given amount between two given accounts
### POST "/get-account"
Send in the body the following:
```
accountID string
```
Retrieve the given account details