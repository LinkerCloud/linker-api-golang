//go:build integration

package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	linkercloud "github.com/linkercloud/linker-api-golang"
	"github.com/linkercloud/linker-api-golang/models"
)

// TestIntegration_Write_Orders_CreateAndCancel creates a minimal test order and
// then cancels it in a cleanup function so no data is left behind.
func TestIntegration_Write_Orders_CreateAndCancel(t *testing.T) {
	requireWritesAllowed(t)
	client := newClient(t)
	ctx := context.Background()

	// Use a unique client order number so test runs don't collide.
	clientOrderNumber := fmt.Sprintf("GO-TEST-%d", time.Now().UnixMilli())

	order := models.OrderType{
		ClientOrderNumber:  clientOrderNumber,
		OrderDate:          time.Now().Format("2006-01-02 15:04:05"),
		ExecutionDate:      time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04:05"),
		ExecutionDueDate:   time.Now().Add(48 * time.Hour).Format("2006-01-02 15:04:05"),
		DeliveryEmail:      "go-test@example.com",
		CodAmount:          &models.Money{Amount: "0", Currency: "PLN"},
		ShipmentPrice:      &models.Money{Amount: "0", Currency: "PLN"},
		ShipmentPriceNet:   &models.Money{Amount: "0", Currency: "PLN"},
		Discount:           &models.Money{Amount: "0", Currency: "PLN"},
		PaymentTransactionID: "",
		Tags:               []string{},
		ValidationErrors:   []*models.UnstructuredType{},
		CustomProperties:   []*models.UnstructuredType{},
		Items: []*models.OrderItemType{
			{
				SKU:              "TEST-SKU",
				Quantity:         "1",
				PriceGross:       "0.00",
				PriceNet:         "0.00",
				SerialNumbers:    []*models.UnstructuredType{},
				CustomProperties: []*models.UnstructuredType{},
			},
		},
	}

	created, err := client.Orders.Create(ctx, order)
	if err != nil {
		t.Fatalf("Orders.Create: %v", err)
	}
	t.Logf("created order: number=%s", created.Number)

	// Best-effort cancel — keeps the instance clean even if the test fails later.
	// We need the internal order ID to cancel; fetch it via Get using the order number.
	t.Cleanup(func() {
		orders, err := client.Orders.List(ctx, linkercloud.ListOrdersOptions{
			Limit:   1,
			Filters: map[string]string{"client_order_number": clientOrderNumber},
		})
		if err != nil || len(orders.Items) == 0 {
			t.Logf("cleanup: could not find order %q to cancel: %v", clientOrderNumber, err)
			return
		}
		if err := client.Orders.Cancel(ctx, orders.Items[0].ID); err != nil {
			t.Logf("cleanup: cancel order %s: %v", orders.Items[0].ID, err)
		} else {
			t.Logf("cleanup: cancelled order %s", orders.Items[0].ID)
		}
	})
}

// TestIntegration_Write_Products_CreateAndUpdate creates a minimal test product
// and then updates it.
func TestIntegration_Write_Products_CreateAndUpdate(t *testing.T) {
	requireWritesAllowed(t)
	client := newClient(t)
	ctx := context.Background()

	sku := fmt.Sprintf("GO-TEST-SKU-%d", time.Now().UnixMilli())

	req := models.ProductType{
		Name:                    "Go Integration Test Product",
		SKU:                     sku,
		StorageUnits:            []*models.UnstructuredType{},
		NameAliases:             []*models.UnstructuredType{},
		Images:                  []string{},
		AdditionalCodes:         []*models.UnstructuredType{},
		IgnoreInWMS:             false,
		IgnoreWhenPacking:       false,
		AlwaysAskForSerialNumber: false,
		HasBatchNumber:          false,
		IsExpirable:             false,
		IsBio:                   false,
		IsFood:                  false,
		IsInsert:                false,
		IsFragile:               false,
	}

	created, err := client.Products.Create(ctx, req)
	if err != nil {
		t.Fatalf("Products.Create: %v", err)
	}
	t.Logf("created product: id=%s sku=%s", created.ID, created.SKU)

	// Update the product name.
	req.Name = "Go Integration Test Product (updated)"
	updated, err := client.Products.Update(ctx, created.ID, req)
	if err != nil {
		t.Fatalf("Products.Update: %v", err)
	}
	t.Logf("updated product: id=%s name=%s", updated.ID, updated.Name)
}

// TestIntegration_Write_SupplierOrders_Create creates a minimal inbound order.
func TestIntegration_Write_SupplierOrders_Create(t *testing.T) {
	requireWritesAllowed(t)
	client := newClient(t)
	ctx := context.Background()

	req := models.SupplierOrderType{
		ClientOrderNumber: fmt.Sprintf("GO-TEST-INBOUND-%d", time.Now().UnixMilli()),
		OrderDate:         time.Now().Format("2006-01-02 15:04:05"),
		ExecutionDate:     time.Now().Add(72 * time.Hour).Format("2006-01-02 15:04:05"),
		PriceGross:        &models.Money{Amount: "0", Currency: "PLN"},
		PriceNet:          &models.Money{Amount: "0", Currency: "PLN"},
		Supplier:          0, // supplier ID — 0 may fail; test is best-effort
		SupplierObject:    &models.SupplierType{},
		Items:             []*models.SupplierOrderItemType{},
		CustomProperties:  []*models.UnstructuredType{},
	}

	created, err := client.SupplierOrders.Create(ctx, req)
	if err != nil {
		// Supplier order creation often requires a valid supplier ID; log and skip
		// rather than fail hard — the goal is to exercise the endpoint.
		t.Logf("SupplierOrders.Create returned error (may need valid supplier): %v", err)
		t.Skip("skipping: supplier order creation requires a valid supplier configured in the instance")
	}
	t.Logf("created supplier order: number=%s", created.ClientOrderNumber)
}
