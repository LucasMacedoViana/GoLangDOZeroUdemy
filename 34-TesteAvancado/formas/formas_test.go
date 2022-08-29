package formas

import (
	"math"
	"testing"
)

func testeArea(t *testing.T)  {
	t.Run("Retâmgulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.area()

		if areaEsperada != areaRecebida{
			t.Fatalf("Area recebida %f  é diferente da area esperada %f", areaRecebida, areaEsperada)
		}
	}) //é usado quando quero separa as funções

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.area()
		if areaEsperada != areaRecebida{
			t.Fatalf("Area recebida %f  é diferente da area esperada %f", areaRecebida, areaEsperada)
		}
	})
}
