package main

func TestSoma(t *testing.T){
	total := Soma(15,15)
	if total != 30 {
		t.Errorf("Error na soma")
	}
}