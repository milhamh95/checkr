# checker

checker is a checkout service that will support different promotions with the given inventory

## Prerequisite

- Go min version >= 1.17

## Product

| SKU | Name | Price | Inventory Qty  |

|--|---|---|---|
| 120P90 | Google Home |  $49.99 | 10 |
| 43N23P | Macbook Pro | $5399.99 | 5 |
| A304SD | Alexa Speaker  | $109.50  | 10 |
| 234234 | Raspberry Pi B  | $30 | 2  |

## Note

- Each sale of a MacBook Pro comes with a free Raspberry Pi B

- Buy 3 Google Homes for the price of 2

- Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa speakers

## Graphql Schema

To find graphql schema, please check

`/graph/schema.graphqls`

There are 2 main mutations

### scanItems

- scanItems is a mutation to insert items to cart

```
type Mutation {
  scanItems(input: ScanItems!): ScanItemsResult!
}
```

- Example Request

```
mutation {
  scanItems(input: {
    items: [
      {
        sku: "120P90",
        quantity: 5,
      },
    ]
  }) {
    message
  }
}
```

### checkout

- checkout is a mutation to finish checkout cart items

```
type Mutation {
  checkout: CheckoutResult!
}
```

- Example Request

```
mutation{
  checkout() {
   totalPrice
    message
    items {
      sku
      name
      price
      discountedPrice
      quantity
      freeQuantity
    }
  }
}
```

## How To Run

### Dev

- Please Run

```
make dev
```

- Open graphql playground in browser

```
http://localhost:8080/
```

### Prod

Please run

```
make build // to build checker binary
make start // to start checker binary
```

- Open graphql playground in browser

```
http://localhost:8080/
```

## Unit Test

To run unit test, please run

```
make unittest
```
