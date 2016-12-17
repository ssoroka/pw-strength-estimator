a toy program that tells you an estimate of how many bits of entropy are in a password.

It uses an internal markov chain built from password data of 10 million leaked passwords.

You can probably find this file elsewhere on the internet. I'd include it in this repo, but it's rather large. Ping me if you want the original password list.

Using the compressed probability space of the markov chain, we store information about the 10 million passwords in a file that's only 207161 bytes, and in JSON at that. It could be reduced quite a bit further if it was stored in a binary format(70%?), but that's not really the point; it's small enough to ship around easily now. The point is that a small password file in the format of a markov chain can tell you roughly how statistically common your particular password is. For example:

Let's say my password is "password". The markov chain looks up the letter "p" and checks to see how often it's followed by the letter "a". Let's say our markov chain says 0.25; 25% of the time it's followed by an "a". I roughly assume this means this "a" is only providing 4 options, thus log2(1/0.25) = 2 bits of entropy. Now this is only a toy and not to be taken too seriously, but by this calculation any password providing less than 60 bits of entropy is generally regarded as a poor password, and anything over something around 60-70 bits can probably be regarded as a "good" password.

Here's an example run that spits out a bunch of passwords scored against the markov-built 10-million password database:

```
$ time go run main.go
password has an estimated 25.640202 bits of entropy.
asdgawegwmgaf has an estimated 64.041785 bits of entropy.
00000000 has an estimated 18.729116 bits of entropy.
12345678 has an estimated 18.953390 bits of entropy.
shotokan has an estimated 24.969701 bits of entropy.
fubgbxna has an estimated 38.991396 bits of entropy.
sometimes I eat cabbage has an estimated 111.162955 bits of entropy.
started has an estimated 20.616509 bits of entropy.
october has an estimated 24.028792 bits of entropy.
octanuary has an estimated 35.430650 bits of entropy.
fishsticks has an estimated 34.901189 bits of entropy.
cucumbers has an estimated 33.573153 bits of entropy.
chesterfield has an estimated 38.467317 bits of entropy.
46I2skN/ has an estimated 51.665273 bits of entropy.
LXu^f4h57 has an estimated 66.661715 bits of entropy.
VztaJ055~ has an estimated 61.560830 bits of entropy.
VztaJ055* has an estimated 55.411082 bits of entropy.
s,r9BzN94e has an estimated 65.388548 bits of entropy.
3Scx2yoW8b^ has an estimated 85.311955 bits of entropy.
!!!!!!! has an estimated 11.755503 bits of entropy.
!!!!!!!! has an estimated 13.714753 bits of entropy.
!!!!!!!!! has an estimated 15.674003 bits of entropy.
!!!!!!!!!! has an estimated 17.633254 bits of entropy.
&&&&&&& has an estimated 25.241774 bits of entropy.
&&&&&&&& has an estimated 29.448736 bits of entropy.
&&&&&&&&& has an estimated 33.655699 bits of entropy.
&&&&&&&&&& has an estimated 37.862661 bits of entropy.
abdominohysterectomy has a score of 79.947183

real	0m0.623s
user	0m0.227s
sys	0m0.243s
```

Interestingly, It makes reasonable predictions of the password quality, with the exception that it doesn't actually know whether or not a word is in a dictionary. Keep in mind it wasn't trained with a dictionary, it was trained with a password list of 10 million leaked passwords.. so passwords real people tend to use, and is as much a reflection of their poor password choices as anything. That tends to work quite well when looking for what to avoid when choosing a password, as that's exactly the type of thing one would guess first.

Cheers.
