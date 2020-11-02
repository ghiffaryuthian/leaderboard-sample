package leaderboard

// User is a struct of leaderboard member detail
type User struct {
	Username string  `json:"username"`
	Score    float64 `json:"score"`
	Rank     int64   `json:"rank"`
}
