package books

type CreatBookRequest struct {
  AuthorID string `json:"authorId"`
  Title string `json:"title"`
  Rating float32 `json:"rating"`
  Description string `json:"description"`
}