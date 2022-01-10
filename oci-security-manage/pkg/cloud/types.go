package cloud

type resource struct {
	name    string
	profile string
	object  string // Security List or NSG, can't use type as it's a keyword
	ocid    string
	id      string
	region  string
}
