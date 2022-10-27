# Pumpkin Stand

## The Problem

Rating: Easy

Flavor Text:
```
This time of the year, we host our big festival and the one who craves the pumpkin faster and make it as scary as possible, gets an amazing prize! Be fast and try to crave this hard pumpkin!
```

Attachments : [pwn_pumpkin_stand.zip](./pwn_pumpkin_stand.zip)

## Solution


break open the provided file:

```
00000b9e  {
00000baf      void* fsbase;
00000baf      int64_t var_10 = *(int64_t*)((char*)fsbase + 0x28);
00000bb5      setup();
00000bba      banner();
00000bbf      int16_t var_54 = 0;
00000bc5      int16_t var_52 = 0;
00000bcb      while (true)
00000bcb      {
00000bcb          menu();
00000be3          __isoc99_scanf(&data_132b, &var_54);
00000bf4          printf("\nHow many do you want?\n\n>> ");
00000c0c          __isoc99_scanf(&data_132b, &var_52);
00000c18          if (var_52 <= 0)
00000c15          {
00000c2d              printf("%s\n[-] You cannot buy less than…", "\x1b[1;31m");
00000c1a          }
00000c62          else
00000c62          {
00000c62              uint64_t rcx_2 = ((uint64_t)(((uint32_t)pumpcoins) - (((uint32_t)var_52) * *(int32_t*)((((int64_t)var_54) << 2) + &values))));
00000c66              pumpcoins = rcx_2;
00000c77              if (pumpcoins < 0)
00000c74              {
00000c9d                  printf("\nCurrent pumpcoins: [%s%d%s]\n\n", "\x1b[1;33m", ((uint64_t)((int32_t)pumpcoins)), "\x1b[1;34m");
00000cbc                  printf("%s\n[-] Not enough pumpcoins for…", "\x1b[1;31m", "\x1b[1;34m");
00000ca2              }
00000cce              else if (var_54 == 1)
00000cca              {
00000cf4                  printf("\nCurrent pumpcoins: [%s%d%s]\n\n", "\x1b[1;33m", ((uint64_t)((int32_t)pumpcoins)), "\x1b[1;34m");
00000d00                  puts("\nGood luck crafting this huge p…");
00000cf9              }
00000d15              else
00000d15              {
00000d15                  if (pumpcoins > 0x270e)
00000d11                  {
00000d15                      break;
00000d15                  }
00000dd7                  printf("%s\n[-] Not enough pumpcoins for…", "\x1b[1;31m", "\x1b[1;34m", rcx_2);
00000dd7              }
00000dd7          }
00000dd7      }
00000d1b      int64_t var_48 = 0;
00000d23      int64_t var_40 = 0;
00000d2b      int64_t var_38 = 0;
00000d33      int64_t var_30 = 0;
00000d3b      int64_t var_28 = 0;
00000d43      int64_t var_20 = 0;
00000d59      FILE* rax_25 = fopen("./flag.txt", &data_13e6);
00000d67      if (rax_25 != 0)
00000d62      {
00000d8f          fgets(&var_48, 0x30, rax_25);
00000dae          printf("%s\nCongratulations, here is the…", "\x1b[1;32m", &var_48);
00000db8          exit(0x16);
00000db8          /* no return */
00000db8      }
00000d70      puts("Error opening flag.txt, please c…");
00000d7a      exit(1);
00000d7a      /* no return */
00000d7a  }

```

we need this to fire

FILE* rax_25 = fopen("./flag.txt", &data_13e6);


bit the loop is while(true)

we have options 1 and 2 but the catch one looks for:

00000d15                  if (pumpcoins > 0x270e)
00000d11                  {
00000d15                      break;
00000d15                  }


Current pumpcoins: [1337]

Items: 

1. Shovel  (1337 p.c.)
2. Laser   (9999 p.c.)

>> 4

How many do you want?

>> 4

Current pumpcoins: [-6215]


[-] Not enough pumpcoins for this!


Current pumpcoins: [-6215]

Items: 

1. Shovel  (1337 p.c.)
2. Laser   (9999 p.c.)

>> 4

How many do you want?

>> 400



## Flag
```
HTB{1nt3g3R_0v3rfl0w_101_0r_0v3R_9000!}

```

## Final Notes
