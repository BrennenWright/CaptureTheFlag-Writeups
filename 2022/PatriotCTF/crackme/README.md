# crack me

## The Problem

Flavor Text:
```


```

Attachments : [crackme](crackme)



## Solution

reversed the program and ran ltrace to confirm that the string check is being done against:
```
9{$d 0DJ0 8e,< 7{zf #emx 9Av@ 7f,A .deB }*/g %xBm w=qq
```
the input variable only holds 31chars

I can assume this is intended to validate the flag as a string so something like this? 

```
pct  f{x  xxx  xxx  xxx  xxx  xxx  xxx  xxx  xxx  }
9{$d0DJh;|>h;|>h;
```

trying a few ltraced runs shows that the first part is correct but after the fourth character things get squirly
```
pctf{1aaaaaaaaaaaaaaaaaaaaaaaa}
9{$d0DIN/-z>/-z>/-z>/-z>/-z>/-z>
```


reviewing the reversed code in ghidra yields:

looks like it alters in groups of three chars
and then makes a fourth that combines using:
 tmpChar3 + tmpChar1 * 0x10000 + tmpChar2 * 0x100;
  
 the fourth character ends up being a=> b=? c=@ or decimal cod - 35
 
 breaking the target string into groups of three I had to guess/check my way through numbers, vowels and then common words to work out the flag

 ```
 
9{$d 0DJ0 8e,< 7{zf #emx 9Av@ 7f,A .deB }*/g %xBm w=qq
pct  f{Y  oU_  hav  3_C  r@c  ked  _Me  !xx  89}
                                             9
 
 now I need:
 8e,< 7{zf #emx 9Av@ 7f,A .deB }*/g %xBm
 
 oUr
 8e,O
 
 oUa
 8e,>
 
oUh
8e,E

oU_
8e,<

hav
7{zf

3_C
#emx

```

after a looooobnnnnngg time I worked out the flag. I reallly need to learn how to use angr 


## Flag
```
pctf{YoU_hav3_Cr@cked_Me!6789}
```

## Final Notes
