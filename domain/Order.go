package domain

type Order struct {
	ID            int
	Userid        int
	PaymentMethod string
	TotalQty      int
	Totalprice    int
	Status        int
}
