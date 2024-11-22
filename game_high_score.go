package telegram

// GameHighScore represents one row of the high scores table for a game.
//
// See "GameHighScore" https://core.telegram.org/bots/api#gamehighscore
type GameHighScore struct {
	// (Required) Position in high score table for the game
	Position int `json:"position"`

	// (Required) User
	User User `json:"user"`

	// (Required) Score
	Score int `json:"score"`
}
