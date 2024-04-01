# treasury

### init blockchain
anvil

### create withdraw claim
curl --location --request POST '127.0.0.1:8080/withdraw/claim/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userID":1,
    "value": 5000000000
}'


### approve withdraw
curl --location --request POST '127.0.0.1:8080/withdraw/approve/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "approverID":2,
    "claimID":1
}'

### get withdraw claim
curl --location --request GET 'http://127.0.0.1:8000/withdraw/claim/1' 

