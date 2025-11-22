package main

import (
	"GoBus/dto"
	"GoBus/handler"
	"fmt"
)

func main() {
	req := dto.NewRequest("sidik", "Jakarta")

	resp, err := handler.ProcessTicket(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("=== Harga Tiket ===")
	fmt.Printf("Penumpang : %s\n", resp.Name)
	fmt.Printf("Tujuan   : %s\n", resp.Destination)
	fmt.Printf("Harga    : RP %.2f\n", float64(resp.Price))
	fmt.Println("===================")
}
