package entity

type EntityType string

const (
	ProductComments EntityType = "product_comments"
)

func (e EntityType) String() string {
	return string(e)
}
