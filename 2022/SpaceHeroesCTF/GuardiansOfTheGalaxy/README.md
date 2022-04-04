Guardians of the Galaxy
100


Ronan the Accuser has the Power Stone. Can Starlord find a successful distraction format? (please note that you should only contact the admin if you recieve the error while using the command below)

nc 0.cloud.chals.io 12690

author: GlitchArchetype


downloaded the file. rand cat and strings. it looks to be a simple c program that works with printing the output, likely the one responding on the nc port.

I played with it a little as nothing but ./flag jumped out at me from the file. 

it responds with OH no each time you enter an input and follows it with a copy of your input. 

the flavor text mentioned a format for the distraction. this is likely a input format vuln where its response can be messed with to provide alternate response instead of the text intended.

looked up c program print format vuls and foudn a few things such as: https://www.geeksforgeeks.org/format-string-vulnerability-and-prevention-with-example/

used that to play with %ps and ended up with %p %p %p %p %p %p before the ps started showing up in the printout they look like: 0x252070 in memory

so I swapped the last p for an %s to output it as a string and got: 
shctf{im_distracting_you}