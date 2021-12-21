package entity

type User struct {
  Id string `json:"id"`
  Email string `json:"email"`
  Name string `json:"name"`
  Password string `json:"password"`
}
