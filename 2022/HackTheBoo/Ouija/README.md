# Ouija

## The Problem

Points: 

Rating:

Author:

Flavor Text:
```
You've made contact with a spirit from beyond the grave! Unfortunately, they speak in an ancient tongue of flags, so you can't understand a word. You've enlisted a medium who can translate it, but they like to take their time...
```

Attachments : [rev_ouija.zip](./rev_ouija.zip)



## Solution

```
strings ouija
...
ZLT{Svvafy_kdwwhk_lg_qgmj_ugvw_escwk_al_wskq_lg_ghlaearw_dslwj!}

```

this turns out to be a rot18 cypher using the brute force we see it plaintext

![](./crack.png)

## Flag
```
HTB{Adding_sleeps_to_your_code_makes_it_easy_to_optimize_later!}

```

## Final Notes


Credit on this one goes to bakajiikara for the solve