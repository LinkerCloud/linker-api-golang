# Integration Tests

These tests run against a **real** Linker Cloud API endpoint. They are **not** part of the normal test suite and must be explicitly enabled with the `integration` build tag.

## Prerequisites

- A Linker Cloud instance URL
- A valid API key for that instance

## Read-only tests (safe for production)

By default, the integration suite only performs **read-only** (GET) operations. It is safe to run against a production instance.

```bash
LINKER_BASE_URL=https://api-mycompany.linker.shop \
LINKER_API_KEY=your-api-key \
go test -tags integration ./integration/... -v
```

| Test | Endpoint |
|------|----------|
| `TestIntegration_Orders_List` | `GET /public-api/v1/orders` |
| `TestIntegration_Products_List` | `GET /public-api/v1/products` |
| `TestIntegration_Stock_List` | `GET /public-api/v1/stocks` |
| `TestIntegration_SupplierOrders_List` | `GET /public-api/v1/supplierorders` |
| `TestIntegration_OrderReturns_List` | `GET /public-api/v1/orderreturns` |
| `TestIntegration_Orders_Get_notFound` | `GET /public-api/v1/orders/{id}` (404 check) |

## Write tests (demo/staging only)

Write tests create and modify real data. They are **skipped by default** and only run when `LINKER_ALLOW_WRITES=1` is set. **Do not run these against a production instance.**

Each write test is designed to clean up after itself (e.g. newly created orders are cancelled in a `t.Cleanup` function).

```bash
LINKER_ALLOW_WRITES=1 \
LINKER_BASE_URL=https://api-mycompany.linker.shop \
LINKER_API_KEY=your-api-key \
go test -tags integration ./integration/... -v
```

| Test | Operations |
|------|------------|
| `TestIntegration_Write_Orders_CreateAndCancel` | `POST /public-api/v1/orders` → `PUT /orders/{id}/cancel` |
| `TestIntegration_Write_Products_CreateAndUpdate` | `POST /public-api/v1/products` → `PUT /products/{id}` |
| `TestIntegration_Write_SupplierOrders_Create` | `POST /public-api/v1/supplierorders` |
