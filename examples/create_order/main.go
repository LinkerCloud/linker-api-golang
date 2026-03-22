// create_order demonstrates how to create a new order.
//
// Run with:
//
//	LINKER_BASE_URL=https://api-mycompany.linker.shop \
//	LINKER_API_KEY=your-api-key \
//	go run ./examples/create_order
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	linkercloud "github.com/linkercloud/linker-api-golang"
	"github.com/linkercloud/linker-api-golang/models"
)

func main() {
	baseURL := os.Getenv("LINKER_BASE_URL")
	apiKey := os.Getenv("LINKER_API_KEY")
	if baseURL == "" || apiKey == "" {
		log.Fatal("LINKER_BASE_URL and LINKER_API_KEY environment variables must be set")
	}

	client := linkercloud.New(baseURL, apiKey)
	ctx := context.Background()

	order := models.OrderType{
		ClientOrderNumber: "MY-SHOP-12345",
		PaymentMethod:     "Bank transfer",
		DeliveryRecipient: "Jane Doe",
		DeliveryEmail:     "jane@example.com",
		DeliveryStreet:    "123 Main Street",
		DeliveryCity:      "Warsaw",
		DeliveryPostCode:  "00-001",
		DeliveryCountry:   "PL",
		Items: []*models.OrderItemType{
			{
				SKU:        "PRODUCT-SKU-001",
				Quantity:   "2",
				PriceGross: "49.99",
			},
		},
	}

	created, err := client.Orders.Create(ctx, order)
	if err != nil {
		if linkercloud.IsUnauthorized(err) {
			log.Fatal("invalid API key")
		}
		log.Fatalf("create order: %v", err)
	}

	fmt.Printf("Order created: number=%s\n", created.Number)
}
