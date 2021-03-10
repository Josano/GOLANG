package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste {
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodóvia dos Imigrantes", "Rodóvia"},
		//{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		//{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", 
			tipoDeEnderecoRecebido, cenario.enderecoEsperado)
		}
	}


}

func testQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("O valor é maior")
	}
}