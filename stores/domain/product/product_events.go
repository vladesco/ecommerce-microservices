package product

type ProductAdded struct {
	Product *Product
}

func (event ProductAdded) GetName() string {
	return "[PRODUCT]: Product Added"
}

type ProductRemoved struct {
	Product *Product
}

func (event ProductRemoved) GetName() string {
	return "[PRODUCT]: Product Removed"
}
