#!/usr/bin/python3

import hmac
import math
import os
import secrets
import datetime as dt

# lets gen a random token and test the hashing to try and get the same result?

secret_key = secrets.token_bytes(32) 

transaction = {
    "from_acct": 1111,
    "to_acct": 314159265,
    "num_tuco": 1
  }
count = 1
t = dt.datetime.now()
target = "906d871d19e461560b88396619b20011964b50df847ab126ef5fb505ee86a416"
  
transaction = (str(1111) + str(314159265) + str(1)).encode()

auth_tag = hmac.new(secret_key, transaction, "sha256").hexdigest()

print("transaction: ", transaction)
print("secret_key: ", secret_key)
print("auth_tag: ", auth_tag)


while auth_tag != target:
    secret_key = secrets.token_bytes(32) 
    auth_tag = hmac.new(secret_key, transaction, "sha256").hexdigest()
    count = count + 1
    delta = dt.datetime.now()-t
    if delta.seconds >= 60:
        print("guessed: ", count)
        t = dt.datetime.now()
    # print("secret_key: ", secret_key)
    # print("auth_tag: ", auth_tag)

print("FOUND IT: ", secret_key)
print("auth_tag: ", auth_tag)

