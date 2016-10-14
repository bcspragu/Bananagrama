# Bananagrama

I like giving people gifts for birthdays, but I also like building useless
things. To kill both these proverbial birds with one proverbial stone, I make
useless sites as gifts. This site (eventually) will enable people to make AIs
of their own devices to play out a slightly modified version of Bananagrams
for cash, prizes, and honor.

## The Rules

It looks kind of like normal Bananagrams, with the following modifications:

* There are 144,000 tiles\*, instead of the usual 144
* Each player starts with 1,000\* times as many tiles as they normally would
* Every time a player calls PEEL successfully, every player receives 100 tiles\*\*
* PEELs don't happen instantaneously, but rather in intervals. So, for example,
  say two players PEEL with valid boards in the same xx ms interval, they'll
  both be considered to have peeled. This is to mitigate race conditions/random
  network issues, and so players aren't perpetually turning in boards that are
  missing the letters they're about to receive.
* The winner isn't the person who calls PEEL last. Currently, the
  n<sup>th</sup> PEEL is worth n points, so PEELs later in the game are worth
  more. The game still ends when there aren't enough tiles for everyone.


\* Set by `TileScalingFactor` in `engine/bunch.go`
\*\* Set by `PeelSize` in `connect.go`
// TODO(bsprague): Add the rules and stuff, like 144,000 tiles and whatnot
