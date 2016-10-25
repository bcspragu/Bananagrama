@0xf2768a70df34b76a;

using Go = import "../../../../zombiezen.com/go/capnproto2/std/go.capnp";

$Go.package("potassium");
$Go.import("github.com/bcspragu/Bananagrama/potassium");

using Tiles = List(Text);
# Tiles are just lists of letters

interface Server {
  connect @0 ConnectRequest -> ConnectResponse;
  # A client connecting to the server
}

struct ConnectRequest {
  # A player connects to the server with a username and their player implementation
  name @0 :Text;
  player @1 :Player;
}

struct ConnectResponse {
  # Response the player gets when they connect to the server
  status @0 :Status;
  enum Status {
    success @0;
    nameTooLong @1; # Don't be an ass, I'll probably cap this at like ~20
    nameAlreadyTaken @2; # Either you're already connected, or you're trying to impersonate someone
    gameIsFull @3; # The running game already has 8 players connected to it
    yaDoneGoofed @4; # No idea what happened, but *something* did
  }
  game @1 :Game;
  # Your interface to the game, on success
}

interface Player {
  # Interface that a competitor implements.

  split @0 SplitRequest -> ();
  # The server calls split when it's time to play, and gives each player a bunch of tiles and a list of opponents

  newTile @1 NewTileRequest -> ();
  # The server calls newTile when someone has called PEEL successfully, letter is the new tile you've been given

  dumpNotice @2 DumpNoticeRequest -> ();
  # Your notice that SOMEONE ELSE has dumped. You won't receive this call when you were the...uh...dumper.

  gameOver @3 () -> ();
  # The game is literally over, stop sending me things
}

struct SplitRequest {
  # Notice that the game is starting, with a list of tiles you received, and a 
  tiles @0 :Tiles;
  players @1 :List(Text);
}

struct NewTileRequest {
  letters @0 :Tiles;
  # The new letters, be sure to add these to your pile. The current setting is to return 100 tiles
  peelers @1 :List(Text);
  # Name of the players who peeled, you can be one of them if you just peeled successfully
}

struct DumpNoticeRequest {
  dumped @0 :Text;
  # The tile that was dumped back to the bunch
  received @1 :List(Text);
  # The list of tiles that was taken by the bunch from the dumper
  dumper @2 :Text;
  # The username of the...here it comes...dumper
}

interface Game {
  peel @0 (board :Board) -> PeelResponse;
  # A player sends their completed board for verification, and receives a
  # success response if it was valid, and a reason if it wasn't valid 
  dump @1 (letter :Text) -> DumpResponse;
  # A player sends a letter to dump, and receives three back if it was valid
}

struct PeelResponse {
  status @0 :Status;
  enum Status {
    success @0; # You peeled successfully, you'll receive your new tile via newTile()
    invalidWord @1; # One or more words the player turned in were shit
    detachedBoard @2; # Not all of the player's letters are connected
    notAllLetters @3; # The player isn't using all of their letters
    extraLetters @4; # The player is trying to play with letters they don't have
    invalidBoard @5; # The player's words overlap in ways they can't, like two words with different letters overlapping
    gameNotStarted @6; # The game hasn't started yet
  }

  union {
    valid @1 :Void; # On success, nothing to send back
    invalidWord @2 :Tiles; # The words that aren't actually words
    notAllLetters @3 :Tiles; # The list of letters that weren't used
    extraLetters @4 :Tiles;
    # The list of letters the player didn't have in their pile, but used
    # anyway. If a player tries to toss some Unicode shit in their words,
    # they'll get this error
  }
  # Huge disclosure: Depending on how lazy I am with my implementation, the
  # errors might not include everything wrong with your board, so just send in
  # not-shitty boards and I'll have to do less work.
}

struct DumpResponse {
  status @0 :Status;
  enum Status {
    success @0;
    letterNotInTiles @1; # The player didn't have the tile they tried to return in their hand
    notEnoughTiles @2; # There are less than three tiles in the pot
    malformedRequest @3; # There wasn't exactly one letter in the request
    gameNotStarted @4; # The game hasn't started yet
  }
  letters @1 :Tiles; # The tiles to send back to the player, on success
}

struct Board {
  # A Player's board
  words @0 :List(Word);
}

struct Word {
  # A word on a bananagrams board

  text @0 :Text;
  # The actual word
  orientation @1 :Orientation;
  # Whether the word should be laid out vertically or horizontally

  enum Orientation {
    unknown @0; # Happens if the user forgot to set the orientation
    horizontal @1;
    vertical @2;
  }

  x @2 :UInt32;
  y @3 :UInt32;
  # The x and y coordinates of the first letter in the word. We need 32-bits to
  # be safe, because we're playing with 144,000 tiles
}

# All the types down here are used for persistence in the DB
struct Peel {
  # The representation of players peeling successfully

  validBoards @0 :List(PlayerBoard);
  # The boards successfully submitted along with who submitted them
  struct PlayerBoard {
    player @0 :Text;
    # The player who peeled successfully
    board @1 :Board;
  }

  newTiles @1 :List(PlayerTile);
  # The tiles returned to everyone

  struct PlayerTile {
    player @0 :Text;
    letters @1 :Tiles;
  }
  timestamp @2 :UInt64;
}

struct Dump {
  # The representation of a player dumping
  player @0 :Text;
  # The player who dumped
  dump @1 :Tiles;
  # The tiles that were dumped
  timestamp @2 :UInt64;
}

struct Replay {
  initialTiles @0 :List(TileFreq);
  # The tiles everyone started with
  struct TileFreq {
    player @0 :Text;
    frequency @1 :List(Int32);
    # A list with 26 elements, one for each letter of the alphabet. The value
    # at each index represents how many of that letter the player has in their hand
  }
  peels  @1 :List(Peel);
  dumps @2 :List(Dump);

  startTime @3 :UInt64;
  endTime @4 :UInt64;
  finalScores @5 :List(Score);
  # The tiles returned to everyone
  struct Score {
    player @0 :Text;
    score @1 :UInt32;
  }
}
