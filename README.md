# linker-api-golang

> ⚠️ **Experimental library — use at your own risk.**
> This client is community-maintained and not officially supported by Linker. The Linker Cloud Public API may change without notice. Verify all responses against the [official API documentation](https://api-doc.linker.shop) before use in production.

A Go client library for the [Linker Cloud](https://linkercloud.com) Public API.

## Installation

```bash
go get github.com/linkercloud/linker-api-golang
```

Requires **Go 1.26+**. Zero external dependencies — only the Go standard library.

## Authentication

All API requests require an API key passed as the `apikey` HTTP header. You receive your key from your Linker operator or customer service. The base URL for your instance is `https://api-<CLIENT_NAME>.linker.shop`.

## Quick start

```go
import linkercloud "github.com/linkercloud/linker-api-golang"

client := linkercloud.New(
    "https://api-mycompany.linker.shop",
    "your-api-key",
)

// List orders
orders, err := client.Orders.List(ctx, linkercloud.ListOrdersOptions{
    Limit:   50,
    SortDir: "DESC",
    Filters: map[string]string{
        "order_date":    "01.01.2024",
        "order_date_to": "31.12.2024",
    },
})

// Create an order
created, err := client.Orders.Create(ctx, models.OrderType{
    ClientOrderNumber: "MY-SHOP-001",
    DeliveryEmail:     "customer@example.com",
    // ...
})

// Get stock levels
stock, err := client.Stock.List(ctx, linkercloud.ListStocksOptions{Limit: 100})

// Get shipment labels (returns decoded bytes, ready to write to file)
labels, err := client.Shipments.GetLabels(ctx, linkercloud.GetLabelsOptions{
    OrderID:   "order-id",
    PackageID: "package-id",
    Format:    linkercloud.LabelFormatPDF,
})
```

## Available services

| Field | Resource | Key operations |
|---|---|---|
| `client.Orders` | Orders | List, Get, Create, Update, Patch, Cancel, ConfirmPicking, UpdatePaymentStatus |
| `client.OrderReturns` | Order Returns | List, Get, Create, Accept |
| `client.OrderAttachments` | Order Attachments | Upload, PrintLabels, GetByOrderNumber, GetByID |
| `client.OrderTransitions` | Order Transitions | GetAvailable, Apply |
| `client.Packing` | Packing | GetContainers, GetMaterials |
| `client.Products` | Products | List, Create, Update |
| `client.Stock` | Stock | List, GenerateReport, GetReport, UpdateProductStocks |
| `client.Shipments` | Shipments / Deliveries | Create, CreatePackages, UpdateStatuses, CancelPackages, GetLabel |
| `client.Tracking` | Tracking Numbers | SetTrackingNumber, SetTrackingNumbers |
| `client.Pickups` | Pickups | Create |
| `client.SupplierOrders` | Inbound Orders | List, Get, Create, Update, Confirm |
| `client.Workflows` | Workflows | Get |
| `client.Sequences` | Sequences | Get |
| `client.Webhooks` | Integration Webhooks | Trigger |

## Filtering, sorting and pagination

List endpoints accept a `Filters` map rendered as `filters[key]=value` query parameters. Date filters use `dd.mm.yyyy` format.

```go
orders, err := client.Orders.List(ctx, linkercloud.ListOrdersOptions{
    Limit:   100,
    Offset:  0,
    SortCol: "order_date",
    SortDir: "DESC",
    Filters: map[string]string{
        "order_date":    "01.01.2024",
        "order_date_to": "31.01.2024",
        "order_status":  "Y",
    },
})
```

All list responses are wrapped in a `Page[T]` envelope:

```go
fmt.Printf("%d of %d orders\n", len(orders.Items), orders.RecordsTotal)
for _, o := range orders.Items { ... }
```

`RecordsTotal` is the total count in the database; `RecordsFiltered` is the count after your filters are applied.

## Order statuses

Outbound orders typically progress through these status codes:

| Status | Meaning |
|--------|---------|
| `N` | New — received, not yet processed |
| `V` | Verified — passed validation, ready for warehouse export |
| `M` | Exported to WMS — in warehouse, awaiting picking |
| `W` | Error — waiting for correction |
| `B` | Exported with stock shortage |
| `L` | Label generated — ready to pack |
| `R` | Ready for packing (optional step) |
| `Y` | Completed — ready to ship / shipped |
| `P` | Picked up by courier (optional) |
| `D` | Delivered (optional) |
| `DC` | Delivery cancelled (optional) |
| `A` | Cancelled |

Note that this depends on an exact workflow configuration.

Inbound (supplier) orders use: `N` → `V` → `M` → `Y` (completed/delivery received).

## Error handling

All non-2xx responses are returned as `*linkercloud.APIError`. When the server returns a JSON body with a `"message"` field, it's extracted automatically:

```go
_, err := client.Orders.Get(ctx, id)
if err != nil {
    var apiErr *linkercloud.APIError
    if errors.As(err, &apiErr) {
        fmt.Println(apiErr.StatusCode) // e.g. 404
        fmt.Println(apiErr.Message)    // e.g. "Unable to find Order entity."
        fmt.Println(apiErr.Body)       // full raw response body
    }
}
```

Convenience helpers:

```go
linkercloud.IsNotFound(err)        // HTTP 404
linkercloud.IsUnauthorized(err)    // HTTP 401
linkercloud.IsForbidden(err)       // HTTP 403 — e.g. order status change not enabled
linkercloud.IsConflict(err)        // HTTP 409
linkercloud.IsTooManyRequests(err) // HTTP 429
```

The client automatically retries once on HTTP 429, honouring the server's `Retry-After` header. If the retry also fails, the error is returned normally.

## Async WMS stock report

Generating a full WMS stock report is asynchronous. Trigger it to get a UUID, then poll until ready:

```go
uuid, err := client.Stock.GenerateReport(ctx, linkercloud.GenerateWMSStockReportOptions{})
if err != nil {
    log.Fatal(err)
}

for {
    report, err := client.Stock.GetReport(ctx, uuid, linkercloud.GetReportOptions{})
    if err != nil {
        var apiErr *linkercloud.APIError
        if errors.As(err, &apiErr) && apiErr.StatusCode == 202 {
            time.Sleep(5 * time.Second) // still generating
            continue
        }
        log.Fatal(err)
    }
    // report is ready
    _ = report
    break
}
```

## Client options

```go
client := linkercloud.New(
    baseURL, apiKey,
    linkercloud.WithTimeout(10*time.Second),          // default: 30s (API gateway max)
    linkercloud.WithHTTPClient(myCustomHTTPClient),   // custom transport/proxy
    linkercloud.WithMaxResponseSize(20<<20),          // default: 10 MB
)
```

When using `WithHTTPClient`, the client inherits the provided client's `Timeout` unless `WithTimeout` is also specified.

## Examples

See the [`examples/`](examples/) directory:

- [`examples/list_orders`](examples/list_orders/main.go) — list orders with date filters
- [`examples/create_order`](examples/create_order/main.go) — create a new order
- [`examples/manage_stock`](examples/manage_stock/main.go) — query and update stock levels

## Running tests

```bash
# Unit tests (no network required)
go test ./...

# Read-only integration tests (safe against any instance)
LINKER_BASE_URL=https://api-mycompany.linker.shop \
LINKER_API_KEY=your-api-key \
go test -tags integration ./integration/... -v

# Write integration tests (creates/cancels real data — use a staging instance)
LINKER_ALLOW_WRITES=1 \
LINKER_BASE_URL=https://api-mycompany.linker.shop \
LINKER_API_KEY=your-api-key \
go test -tags integration ./integration/... -v
```

## How it works

The client is a thin HTTP wrapper around the Linker Cloud REST API. It:

1. Injects the `apikey` header on every request via a custom `http.RoundTripper`.
2. Serialises request bodies to JSON and deserialises successful responses back to typed Go structs.
3. Returns `*APIError` (carrying the HTTP status code and raw response body) for any non-2xx response.

Model types in the [`models/`](models/) package are derived from the official Linker Cloud OpenAPI specification and verified against the live API.

## License

[MIT](LICENSE)
