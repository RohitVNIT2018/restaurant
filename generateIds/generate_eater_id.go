package order

var Id = 12000

//simple sequence id generation
func Generate_eater_id() int {
	Id++

	return Id
}
