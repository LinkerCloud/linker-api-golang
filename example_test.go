package linkercloud_test

import (
	"context"
	"fmt"
	"log"
	"time"

	linkercloud "github.com/linkercloud/linker-api-golang"
	"github.com/linkercloud/linker-api-golang/models"
)

func ExampleNew() {
	client := linkercloud.New(
		"https://api-mycompany.linker.shop",
		"your-api-key",
		linkercloud.WithTimeout(10 * time.Second),
	)
	_ = client
}

func ExampleOrdersService_List() {
	client := linkercloud.New("https://api-mycompany.linker.shop", "your-api-key")

	orders, err := client.Orders.List(context.Background(), linkercloud.ListOrdersOptions{
		Limit:   50,
		SortCol: "order_date",
		SortDir: "DESC",
		Filters: map[string]string{
			"order_date":    "01.01.2024",
			"order_date_to": "31.12.2024",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d orders (of %d total)\n", len(orders.Items), orders.RecordsTotal)
}

func ExampleOrdersService_Create() {
	client := linkercloud.New("https://api-mycompany.linker.shop", "your-api-key")

	created, err := client.Orders.Create(context.Background(), models.OrderType{
		ClientOrderNumber: "MY-SHOP-001",
		DeliveryEmail:     "customer@example.com",
		OrderDate:         "2024-01-15 10:00:00",
		ExecutionDate:     "2024-01-16 10:00:00",
		ExecutionDueDate:  "2024-01-17 10:00:00",
		Items: []*models.OrderItemType{
			{SKU: "PRODUCT-001", Quantity: "2", PriceGross: "49.99"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created order: %s\n", created.Number)
}

func ExampleIsNotFound() {
	client := linkercloud.New("https://api-mycompany.linker.shop", "your-api-key")

	_, err := client.Orders.Get(context.Background(), "nonexistent-id")
	if linkercloud.IsNotFound(err) {
		fmt.Println("order not found")
	}
}
