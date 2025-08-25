package domain

type ProductInput struct {
    Name  string  `json:"name"`
    Description string `json:"description"`
    ImgUrl string `json:"img_url"`
    UserId int64 `json:"user_id"`
}