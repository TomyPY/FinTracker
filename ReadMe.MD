# FINANCES ORGANIZATION API REST

## This API will serve to organization all your spends/earns

To allow this will be different USERS with a UNIQUE ID, USERNAME will be unique too and a WALLET.
The WALLET will start always with $0 the idea is that the user can interact with the wallet to increase or decrease the money. This will be reached with TRANSACTIONS, transactions will be structs with a UNIQUE ID, MONEY_AMOUNT and TYPE. TYPE will be BUY or SELL, where BUY will increase the wallet by the amount specified and SELL will decrease the walley by the amount specified

## Features:

* Increase the money with N amount of money
* Decrease the money with N amount of money
* Store the history of transactions

## Endpoints: 

### WALLETS
* GET /wallets
This endpoint obtain all wallets

* GET /wallets/wallet_id
This endpoint obtain a wallet

* GET /wallets/wallet_id/transactions
This endpoint get all transactions of a wallet you can send a query param to specify a transaction

* POST /wallets/wallet_id/add
receive a form to add money to the wallet

* POST /wallets/wallet_id/substract
receive a form to substract money from the wallet

* DELETE /wallets/wallet_id
deletes a specifies wallet

### USERS
* GET /users/user_id
* POST /users

## Entities

WALLET:
- id
- money

WALLET_HISTORY: 
- id
- history

USER:
- id
- username
- wallet_id

TRANSACTION:
- ID
- AMOUNT
- TYPE
