# Super TicTacToe

Inspired by a VSauce short about Super TicTacToe, I decided to make it!
I also decided that this will be my first project written in Go, so it
should be a fun, diverse, and complicated learning experience.

## Roadmap

I have a checklist, that I loosely plan to complete in the following order:

- [x] Make a working 3x3x3x3 grid as required by the game without making
the code look too ugly with 4d indexing all over the place.

- [x] When one of the subgames has won, that board displays as a giant X or O,
rather than the actual contents of the board.

- [x] Create a player/computer turn-based input loop. The computer AI will be 
random choosing locations. All this is text/console based.

- [x] Both parties must conform to the rules of the game, restricting the possible
moves they can make in many situations. Terminate the game when someone wins. 

- [ ] Implement a smarter AI (minimax, alpha-beta pruning, and/or dynamic programming).

- [ ] Transform this into an otherwise equivalent client/server HTTP API based design.

- [ ] Create a graphical frontend.
