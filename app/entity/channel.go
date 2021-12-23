package entity

type Channel struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Description string `json:"discreption"`
  Dm_flag string `json:"dm_flag"`
  Member []int `json:"member"`
}
