package entity

import "time"

// ProductComment は商品コメントのエンティティ
type ProductComment struct {
	ID        uint64    `json:"id"`
	ProductID uint64    `json:"product_id"`
	UserID    uint64    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Rate      uint32    `json:"rate"`
	CreatedAt time.Time `json:"created_at"`
}

// NewProductComment はエンティティのコンストラクタ
func NewProductComment(id, productID, userID uint64, title, content string, rate uint32) *ProductComment {
	return &ProductComment{
		ID:        id + 1, // IDは最新のID+1のデータを設定
		ProductID: productID,
		UserID:    userID,
		Title:     title,
		Content:   content,
		Rate:      rate,
		CreatedAt: time.Now(),
	}
}
