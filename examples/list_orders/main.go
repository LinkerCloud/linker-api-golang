// list_orders demonstrates how to list orders with filtering and pagination.
//
// Run with:
//
//	LINKER_BASE_URL=https://api-mycompany.linker.shop \
//	LINKER_API_KEY=your-api-key \
//	go run ./examples/list_orders
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	linkercloud "github.com/linkercloud/linker-api-golang"
)

func main() {
	baseURL := os.Getenv("LINKER_BASE_URL")
	apiKey := os.Getenv("LINKER_API_KEY")
	if baseURL == "" || apiKey == "" {
		log.Fatal("LINKER_BASE_URL and LINKER_API_KEY environment variables must be set")
	}

	client := linkercloud.New(baseURL, apiKey)

	ctx := context.Background()

	// List the 10 most-recently updated orders.
	orders, err := client.Orders.List(ctx, linkercloud.ListOrdersOptions{
		Limit:   10,
		SortCol: "order_date",
		SortDir: "DESC",
		// Filter by a specific date range using the API's filter syntax.
		Filters: map[string]string{
			"order_date":    "01.01.2024",
			"order_date_to": "31.12.2024",
		},
	})
	if err != nil {
		if linkercloud.IsUnauthorized(err) {
			log.Fatal("invalid API key")
		}
		log.Fatalf("list orders: %v", err)
	}

	fmt.Printf("Found %d / %d orders:\n", len(orders.Items), orders.RecordsTotal)
	for _, o := range orders.Items {
		fmt.Printf("  - [%s] %s  status=%s\n", o.ID, o.Number, o.OrderStatus)
	}
}
