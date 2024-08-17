package domain

type CustomerDomain struct {
	Name string `json:"username" xml:"username"`
	Age  int    `json:"age" xml:"age"`
}

func GetCustomers() []CustomerDomain {
	customers := []CustomerDomain{
		{Name: "Amir", Age: 20},
		{Name: "Mahyar", Age: 22},
		{Name: "Nil", Age: 22},
	}
	return customers
}
