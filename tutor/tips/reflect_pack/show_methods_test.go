package reflect_pack

import "testing"

func TestUser_Hello(t *testing.T) {
	u := User{1, "zs", 20}
	Poni(u)
}
