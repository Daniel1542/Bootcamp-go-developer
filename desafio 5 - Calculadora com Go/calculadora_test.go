package main

import "testing"

func TestSoma(t *testing.T) { // ShouldSumCorrect
	teste := soma(1, 2, 5)
	resultado := 8

	if teste != resultado { 
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSoma2(t *testing.T) { // ShouldSumIncorrect
	teste := soma(1, 2, 5)
	resultado := 7 

	if teste != resultado { 
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestMult(t *testing.T) { // ShouldMultCorrect
	teste := multiplica(10, 3)
	resultado := 30
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestMult2(t *testing.T) { // ShouldMultIncorrect
	teste := multiplica(10, 3)
	resultado := 2560
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSub(t *testing.T) { // ShouldSubCorrect
	teste := subtrai(2, 2)
	resultado := 0
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func TestSub2(t *testing.T) { // ShouldSubIncorrect
	teste := subtrai(2, 2)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func TestDiv(t *testing.T) { // ShouldDivCorrect
	teste := divide(100)
	resultado := 50
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
func TestDiv2(t *testing.T) { // ShouldDivIncorrect
	teste := divide(100)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
}
}
	