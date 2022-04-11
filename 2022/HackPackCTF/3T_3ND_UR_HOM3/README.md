# 3T 3ND UR HOM3

## The Problem

Points: 449

Rating: medium

Author: kiwi

Flavor Text:
```
ET escaped from capture! He seems to have lived on until 2022, and he followed you back in his own time machine. He's upgraded his security and gone back in time, and changed the flip phone to a Pixel. Maybe that's why the previous challenge didn't contain an actual invasion date?

Now is your turn to find the actual hidden date that ET plans to conquer the world, so we can be prepared.
```

Attachments : [totallynotinvasion.apk](totallynotinvasion.apk)

hint
```
You can verify that you got the right flag from within the app itself. This should help you avoid red herrings!
```

## Solution

downloaded the source

installed on my phone to check it out. 

it takes a flag input and responds yes or no

ghidra didnt work well here so I switched over to jagx and found this key check:

```
if (this.a.getText().toString().equals(new qg().c(pg.b().replaceAll("[^\\d]", "").substring(0, 15), string))) {
            str = "Invasion Date Verified.\n Welcome, Zreck.";
        } else {
            Toast.makeText(this, "You have been scanned. Human is not in the aliens file.", 1).show();
            str = "This incident will be reported.";
        }
```


so we need to reverse the:
qg().c(pg.b().replaceAll("[^\\d]", "").substring(0, 15), string)

lets work out what qg().c does

```

    public String c(String str, String str2) {
        try {
            return new String(a(str.getBytes("UTF-8"), Base64.decode(str2, 0)));
        } catch (Exception e) {
            e.printStackTrace();
            return str2;
        }
    }

```

## Flag

not found
