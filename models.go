type customer struct {
	Name string
	Age  int
}

type product struct {
	Name string
}

type order struct {
	Customer customer
	Item     product
}