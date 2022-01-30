# WebSocket Messages

## App Flow
1. Connect to app
2. Upgrade to WebSocket connection
3. Get identification number
4. Enter identification number
5. Connect to game
6. Put in X/O
7. Allowable play area changes
8. Other user puts in X/O
9.  A user wins the game
10. A user wishes to play again -> unimportant, works without this
11. A user leaves the game -> unimportant, works without this

## Get ID Messages
### Frontend to Backend
```JSON
{ "command": "requestId" }
```
### Backend to Frontend
```JSON
{ "command": "respondId", "id": number }
```

## Send ID Messages
### Frontend to Backend
```JSON
{ "command": "joinGame", "id": number }
```
### Backend to Frontend
If the ID is found:
```JSON
{ 
    "command": "successIdFound", 
    "message": "ID found, game beginning"
}
```
If the ID is not found:
```JSON
{
    "command": "errorIdNotFound",
    "message": "User with given ID not found"
}
```

## Game Messages
_Payloads defined in TypeScript for brevity, extrapolate from payload constant_
### Frontend to Backend
```typescript
enum Player {
    X = 'X',
    O = 'O',
    NONE = ''
}

interface Coordinate {
    x: number
    y: number
}

interface newMove {
    coordinates: Coordinate[]
    player: Player
}

const payload = '{ "command": "makeMove", "move": /* Insert valid stringified JSON newMove here */ }'
```

### Backend to Frontend
```typescript
enum Player {
    X = 'X',
    O = 'O',
    NONE = ''
}

interface BoardState {
    winner: Player
    isPlayable: boolean
    innerContents: BoardState[][] | Player [][]
}

interface GameState {
    winner: Player
    turn: Player
    board: BoardState
}

const payload = '{ "command": "updateGameState",  "gameState": /* Insert valid stringified JSON GameState here */ }'
```

## Play Again Messages
### Frontend to Backend
```JSON
{ "command": "playAgain" }
```
### Backend to Frontend
If the other player responds in kind:
```JSON
{
    "command": "successPlayAgain",
    "message": "Game beginning"
}
```
If the other player does not:
```JSON
{
    "command": "errorPlayAgain",
    "message": "The other player does not wish to play again"
}
```

## Leave Game Messages
### Frontend to Backend
```JSON
{ "command", "leaveGame" }
```
### Backend to Frontend
```JSON
{
    "command": "successLeaveGame", 
    "message": "The game was left successfully"
}
```