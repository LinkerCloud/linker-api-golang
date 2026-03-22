//go:build integration

// Package integration contains tests that run against a real Linker Cloud API
// endpoint. They are excluded from normal test runs and must be opted in with
// the "integration" build tag.
//
// Run with:
//
//	LINKER_BASE_URL=https://api-mycompany.linker.shop \
//	LINKER_API_KEY=your-api-key \
//	go test -tags integration ./integration/... -v
package integration

import (
	"context"
	"testing"

	linkercloud "github.com/linkercloud/linker-api-golang"
)

func TestIntegration_Orders_List(t *testing.T) {
	client := newClient(t)
	ctx := context.Background()

	orders, err := client.Orders.List(ctx, linkercloud.ListOrdersOptions{
		Limit: 5,
	})
	if err != nil {
		t.Fatalf("Orders.List: %v", err)
	}

	t.Logf("received %d / %d orders", len(orders.Items), orders.RecordsTotal)
	for _, o := range orders.Items {
		t.Logf("  order: id=%s number=%s status=%s", o.ID, o.Number, o.OrderStatus)
	}
}

func TestIntegration_Products_List(t *testing.T) {
	client := newClient(t)
	ctx := context.Background()

	products, err := client.Products.List(ctx, linkercloud.ListProductsOptions{
		Limit: 5,
	})
	if err != nil {
		t.Fatalf("Products.List: %v", err)
	}

	t.Logf("received %d / %d products", len(products.Items), products.RecordsTotal)
	for _, p := range products.Items {
		t.Logf("  product: id=%s sku=%s", p.ID, p.SKU)
	}
}

func TestIntegration_Stock_List(t *testing.T) {
	client := newClient(t)
	ctx := context.Background()

	stock, err := client.Stock.List(ctx, linkercloud.ListStocksOptions{
		Limit: 5,
	})
	if err != nil {
		t.Fatalf("Stock.List: %v", err)
	}

	total := int64(0)
	if stock.TotalCount != nil {
		total = *stock.TotalCount
	}
	t.Logf("total stock records: %d", total)
	for _, item := range stock.Items {
		free := int64(0)
		if item.Free != nil {
			free = *item.Free
		}
		t.Logf("  sku=%s free=%d", item.SKU, free)
	}
}

func TestIntegration_SupplierOrders_List(t *testing.T) {
	client := newClient(t)
	ctx := context.Background()

	orders, err := client.SupplierOrders.List(ctx, linkercloud.ListSupplierOrdersOptions{
		Limit: 5,
	})
	if err != nil {
		t.Fatalf("SupplierOrders.List: %v", err)
	}

	t.Logf("received %d / %d supplier orders", len(orders.Items), orders.RecordsTotal)
}

func TestIntegration_OrderReturns_List(t *testing.T) {
	client := newClient(t)
	ctx := context.Background()

	returns, err := client.OrderReturns.List(ctx, linkercloud.ListOrderReturnsOptions{
		Limit: 5,
	})
	if err != nil {
		t.Fatalf("OrderReturns.List: %v", err)
	}

	t.Logf("received %d / %d order returns", len(returns.Items), returns.RecordsTotal)
}

func TestIntegration_Orders_Get_notFound(t *testing.T) {
	client := newClient(t)
	ctx := context.Background()

	_, err := client.Orders.Get(ctx, "nonexistent-order-id-xxxxxxxxxxx")
	if err == nil {
		t.Fatal("expected error for non-existent order, got nil")
	}
	if !linkercloud.IsNotFound(err) {
		t.Logf("got error (may not be 404 on this instance): %v", err)
	}
}
