// manage_stock demonstrates how to query stock levels and update product stocks.
//
// Run with:
//
//	LINKER_BASE_URL=https://api-mycompany.linker.shop \
//	LINKER_API_KEY=your-api-key \
//	go run ./examples/manage_stock
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

	// --- Query current stock levels ---
	stock, err := client.Stock.List(ctx, linkercloud.ListStocksOptions{
		Limit: 25,
	})
	if err != nil {
		log.Fatalf("list stock: %v", err)
	}

	total := int64(0)
	if stock.TotalCount != nil {
		total = *stock.TotalCount
	}
	fmt.Printf("Total stock records: %d\n", total)
	for _, item := range stock.Items {
		free := int64(0)
		if item.Free != nil {
			free = *item.Free
		}
		fmt.Printf("  SKU=%-20s  free=%d\n", item.SKU, free)
	}

	// --- Update stock for specific products ---
	updateReq := models.StockUpdateRequest{
		Items: []*models.StockUpdateItemRequest{
			{
				SKU:           "PRODUCT-SKU-001",
				TotalQuantity: 100,
			},
		},
	}

	if err := client.Stock.UpdateProductStocks(ctx, updateReq); err != nil {
		log.Fatalf("update stock: %v", err)
	}
	fmt.Println("Stock updated successfully.")
}
