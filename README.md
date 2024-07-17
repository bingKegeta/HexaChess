# HexaChess

HexaChess is a web-based game that brings a unique twist to the classic game of chess by introducing a hexagonal board following the rules of [Glinski's Hexagonal Chess](https://en.wikipedia.org/wiki/Hexagonal_chess#Gli%C5%84ski's_hexagonal_chess). This project aims to provide an engaging and strategic gameplay experience with a modern web application interface.

## Table of Contents

- [HexaChess](#hexachess)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Features](#features)
  - [Technologies Used](#technologies-used)
  - [Features Implementation](#features-implementation)
    - [Backend:](#backend)

## Overview

HexaChess is a personal project that reimagines traditional chess on a hexagonal board. This game is built using React with TypeScript for the frontend and Go for the backend. The goal is to create a seamless and intuitive user experience while offering a challenging and enjoyable game of chess.

## Features

- **Hexagonal Chessboard**: A unique hexagonal chessboard layout that adds a new layer of strategy to the game.
- **Modern UI**: A clean and responsive user interface built with React and TailwindCSS (planned).
- **Backend with Go**: A robust backend powered by Go to handle game logic and player interactions (planned).
- **Real-time Gameplay**: Interactive and real-time gameplay experience (planned).

## Technologies Used

- **Frontend**: React, TypeScript, TailwindCSS (planned)
- **Backend**: Go
- **Database**: SQLite (planned)

> P.S. I do have plans on containerizing the entire project using Nix/Docker but will start that after the MVP is finished first :D

## Features Implementation

Key:
- ✅ : Completed with basic testing
- 🟡 : Currently in progress/testing
- 🔵 : Will be worked on

### Backend:
| Feature                                          | Status | 
|--------------------------------------------------|--------|
| Create Skeletons for the pieces                  |   ✅   |
| Create Board Skeleton                            |   ✅   |
| Create Board Coordinate System                   |   ✅   |
| Add Board Boundary Checks                        |   ✅   |
| Set the pieces on the Board                      |   ✅   |
| Capturing Functions                              |   🟡   |
| Piece Movement Functions                         |   🟡   |
| Successfully start a new game                    |   🟡   |
| Rotate Turns                                     |   🔵   |
| Sucessfully detect checks                        |   🔵   |
| Detect Checkmates                                |   🔵   |
| Halt the game on Checkmate/Detect Winner         |   🔵   |
| Real-time Gameplay                               |   🔵   |



More information will be published as project continues!