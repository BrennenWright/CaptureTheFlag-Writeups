# Curious?

## The Problem

Points: 100
Rating: Easy

Flavor Text:
```
I've been thinking about going on vacation recently and need to know where off Earth I can find these dunes. What sol was this image taken on?

Flag format: shctf{SOL_xxx}
```

Attachments : [dunes.PNG](dunes.PNG)

## Solution
downloaded dunes.PNG

the image file doesn't include anything interesting

image searched the thing and found a site doing image work for soil detection but cant find the original for capture dates or anything
![](reversesearch.png)


the other challenges of this os int section are solved using web searching

I figure this has something to do with mars as its a SpaceX crowd but the flags:
shctf{SOL_MARS} or shctf{SOL_MAR} doesn't work

the flag seems to need three characters and mentioned SOL like in our solar system or star... but it could mean something like day or month? 

the name of the challenge Curious made me think Curiosity (rover) which is a mars mission
![](nasa.png)

searched for curiosity and got the rover site on wikipedia
searched it for SOL

found a reference to  sols (Martian days). 

hit the nasa mission site and found more mentions to tracking the mission in sols. this means we are looking for the numerical mission date XXX the photo was taken on

found the https://mars.nasa.gov/msl/mission-updates/?page=2 mars mission updates blog and it posts pictures by dates :)


found the images and raw images listed https://mars.nasa.gov/msl/multimedia/images/?page=0&per_page=25&order=pub_date+desc&search=&category=51%3A176&fancybox=true&url_suffix=%3Fsite%3Dmsl


found nothing in the blog and image lib, managed to find the explore with curiosity on the project site. found a few early high point view dates that had three digit sols. 

dug into "dingo gap" and realized the image is flipped. this made the search fail? 
ill try to swap it and do the search again for kicks no luck
![](rovertracker.png)


## Flag

the mission tracker has the SOL 533 listed in the details to the right.

```
shctf{SOL_533}
```

## Final Notes
fun little intel puzzle. 