checkersd tx checkers create-game $alice $bob --from $alice --gas auto
checkersd tx checkers play-move 1 1 2 2 3 --from $alice
checkersd query checkers show-stored-game 1
