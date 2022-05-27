# Rock and Roll

## The Problem

Points: 

Rating: Beginner

Author:

Flavor Text:
```
Beginner
The website is blank! I can't see anything. Can you help me find what's hidden?


chal2.pctf.competitivecyber.club:49358
```

Attachments : []()



## Solution

the site quickly redirects to the rick roll video very quickly. So fast you cant check the source for the page reliably. 

soo a quick curl will pull it down and stop the redirection

curl -v chal2.pctf.competitivecyber.club:49358

results in:
![](flag.png)

## Flag

```
PCTF{r1Ck_D0wn_7h3_r0ll}

```

## Final Notes
