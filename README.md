# Bananagrama

This is a web-based implementation of Bananagrams. This repo used to hold a
_very_ different codebase which was for Bananagrams between RPC-slinging AIs,
but the codebase has been repurposed to build a web version of Bananagrams. The
purpose is mainly to facilitate my yearly Christmas gift to my friends: a
drunken web-game-based tournament for cash, prizes, fame, and glory.

## How It Works

1. Pick a name
1. Create or join a game that isn't in-progress or finished.
1. The game will start when the game creator hits 'Start Game'.

## How to Play

It's basically normal Bananagrams:

* The tiles at the bottom are the tiles in your hand.
* Place words by typing them out and hitting `Enter`.
* The game will suggest places to play your words, use the left and right
  arrows to cycle though suggestions, and `Esc` to cancel.
* Click a tile on the board to place it back in your hand.
* Double-click a tile in your hand to dump it back to the bunch (in exchange
  for three new tiles from the bunch).
* To place a word manually, click the tile in your hand, and click the empty
  space on the board where it should be placed.
* When a player PEELS (i.e. arranges all the tiles in their hand into a valid
  board), all players will receive a new tile from the bunch. This is
  represented by a red flash for the non-PEELing players.

## Modifications from traditional Bananagrams

* Like in real Bananagrams, you can put invalid words on your board. Unlike in
  real Bananagrams, you can't PEEL with an invalid board. This means there's no
  validation at the end of the game where other players can scrutinize your
  board.
* If there are more than 8 players, the number of tiles in the bunch increases.
  For every 8 players beyond the initial 8, another 144 tiles will be added to
  the bunch at the start of the game, following the same distribution as the
  initial 144.
* [Planned, but not implemented] Instead of just having a winner and losers,
  there will actually be a ranking of players. The winner will be whoever
  issued the last peel. Remaining players will be ordered by how many times
  they peeled, with ties broken by how many letters they had on their board.
  There are certainly problems with this approach, but I have no plans to fix
  them in the short term.

## Running locally

### Backend

The backend is written in Go and the main binary lives in `cmd/server`. From
the root of the project, run:

```
cd cmd/server && go run main.go --dict=dict.txt
```

`--dict` should specify the path to the dictionary to get the valid word list
from. The dictionary should be formatted as a single, upper-cased word per
line.

NOTE: Since we're just using `go run`, there's currently no support for
live-reloading of the backend, which would currently be annoying anyway
because the DB runs in-memory.

### Frontend

The frontend is written in
[TypeScript](https://www.typescriptlang.org/)/[Vue.js](https://vuejs.org/). I
don't like polluting my local computer with JS development tools, so I run the
frontend in a Docker container.


#### In a Docker container

To run in a Docker container, you'll need an image built similarly to the
following Dockerfile:

```Dockerfile
# Dockerfile

FROM node:alpine
VOLUME /project
WORKDIR /project

RUN apk update
ENV PATH="/.yarn/bin:${PATH}"
RUN mkdir -p /.cache/yarn /.yarn /.config/yarn \
    && touch /root/.profile \
    && touch /.yarnrc \
    && chmod 777 /.cache/yarn /.yarn /.config/yarn /.yarnrc \
    && yarn config set prefix /project
```

Then, from the same directory that it's in, build it with:

```
docker build -t node-env .
```

Now you can use the `frontend/serve.sh` script to run the frontend, with live
reloading.

#### Not in a Docker container

If you have a local JS environment containing Yarn and all that jazz, you
_should_ be able to run:


```
$ cd frontend
$ yarn # to install node_modules
$ yarn serve # to run server
```

But I've never tested this personally.

## Future Work

Check Github issues for current limitations and things I might eventually plan
on fixing. Contributions are welcome!
