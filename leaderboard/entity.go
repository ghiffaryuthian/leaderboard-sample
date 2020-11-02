package leaderboard

// User is a struct of leaderboard member detail
type User struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
	Rank  int64   `json:"rank"`
}
