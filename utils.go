package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Platillo struct {
	ID               int     `json:"id"`
	NombrePlatillo   string  `json:"nombre_platillo"`
	PrecioPorCadaUno float64 `json:"precio_por_cada_uno"`
	Cantidad         int     `json:"cantidad"`
}

type Orden struct {
	Orden []Platillo `json:"orden"`
}

func formatOrder(orderJson string) (string, error) {
	// JSON original

	// Parsear el JSON
	var orden Orden
	if err := json.Unmarshal([]byte(orderJson), &orden); err != nil {
		fmt.Println("ERROR AT PARSING JSON ORDER")
		return "", err
	}

	// Imprimir el detalle de cada platillo y calcular el total
	var output strings.Builder
	var total float64
	for _, platillo := range orden.Orden {
		subtotal := platillo.PrecioPorCadaUno * float64(platillo.Cantidad)
		total += subtotal
		output.WriteString(fmt.Sprintf("- %s (x%d): $%.2f\n", platillo.NombrePlatillo, platillo.Cantidad, subtotal))
	}

	output.WriteString(fmt.Sprintf("\nTotal del pedido: $%.2f\n", total))

	// Obtener la salida como una cadena
	result := output.String()
	return result, nil
}
