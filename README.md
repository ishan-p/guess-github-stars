## Guess the stars - a practice Go CLI game

This is a CLI game in which given a trending public repository on Github, you have to guess the number of stars it will have.

How does the game work:
- You will be optionally asked to choose a language
- There will be 5 rounds in total
- Each round presents you with a new trending repository
- You win the round by guessing the stars within 10% tolerance
- You win the game if you win at least 4 of the 5 rounds
