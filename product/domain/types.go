package domain

type ProductInput struct {
    Name  string  `json:"name"`
    Description string `json:"description"`
    ImgUrl string `json:"img_url"`
    UserId string `json:"user_id"`
}