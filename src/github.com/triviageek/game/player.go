package game

type Player struct {
	Pseudo string `json:"pseudo,omitempty"`
	Game   *Game  `json:"game,omitempty"`
	Score  int    `json:"score,omitempty"`
}
