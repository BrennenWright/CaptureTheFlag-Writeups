# Ghost Wrangler

## The Problem

Rating: easy

Flavor Text:
```

Who you gonna call?

```

Attachments : [rev_ghost_wrangler.zip](rev_ghost_wrangler.zip)



## Solution

opened in binaryninja 

found the flags source text and ran it in cyberchef and this turned up soemthing very close to an xor


the read flag looks to xor so that confirms it. the readflag seems to use the key 0x13
you could jump to this before the close or just xor the the source

```
[GQh{\'f}g wLqjLg{ Lt{#`g&L#uLpgu&Lc\'&g2n
```


the result on this printed the 4 as 04 causing a failed submit. manual removal of those fixed it


## Flag
```
HTB{h4unt3d_by_th3_gh0st5_0f_ctf5_p45t!}
```

## Final Notes
