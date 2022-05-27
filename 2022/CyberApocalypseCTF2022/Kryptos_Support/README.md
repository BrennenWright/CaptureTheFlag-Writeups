# Kryptos Support

## The Problem


Flavor Text:
```
The secret vault used by the Longhir’s planet council, Kryptos, contains some very sensitive state secrets that Virgil and Ramona are after to prove the injustice performed by the commission. Ulysses performed an initial recon at their request and found a support portal for the vault. Can you take a look if you can infiltrate this system?



```




## Solution

I neglected to type up notes as I worked through the process on this one. it was a multistep solve and ended up being a really good excuse to practice my golang solve scripts. 


[solve.go](solve.go)


Step one: We attack the moderator by XSS the service is a ticket input system. The payload triggers the request to the webhook site including the users session cookie

Step Two: We use the cookie to post a password update to the system but flip the user id as it only requires the user id and new password. 

Step three: login as the admin and review the backed for the flag




## Flag
```
HTB{x55_4nd_id0rs_ar3_fun!!}

```

## Final Notes
