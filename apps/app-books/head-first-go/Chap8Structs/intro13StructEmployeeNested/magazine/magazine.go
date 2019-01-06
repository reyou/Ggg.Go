package magazine

// Subscriber is magazine subscriber
type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

// Address is an address
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

// Employee is something cool!
type Employee struct {
	Name    string
	Salary  float64
	HomeAddress Address
}
