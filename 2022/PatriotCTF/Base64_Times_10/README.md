# Base64 Times 10

## The Problem

Points: 498

Rating: 

Author: Necktie (Maxime Bonnaud)

Flavor Text:
```

I heard that rot13 and base64 are easily decoded, but I'm sure that it's just because they are only used once. As such, I have masterfully hidden the flag behind 1000 layers of rot13 and 10 layers of base64. It should be pretty much impossible for anyone to get it now!


```

Attachments : [cipher.txt](cipher.txt)

## Solution

so decode the base64x10

```
cat cipher.txt | base64 --decode | base64 --decode | base64 --decode | base64 --decode | base64 --decode | base64 --decode | base64 --decode | base64 --decode | base64 --decode | base64 --decode
```


## Flag
```
pctf{0bfusc@tion_1s_n0t_3ncrypt10n}
```

## Final Notes
