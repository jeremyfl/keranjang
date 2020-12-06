# Keranjang

Expandana Test

## How to Run

`go run main.go`

## API Documentation

### Get

/api/v1/carts

#### Response

```
{
    "data": [
        {
            "product_id": 1,
            "name": "Macbook Pro M1",
            "color": "Gray"
        }
    ],
    "message": "",
    "status_code": 200
}
```

### Post

/api/v1/carts

#### Request

```
{
    "product_id": 1,
    "name": "Macbook Pro M1",
    "color": "Gray"
}
```

### Response

```
{
    "data": {
        "product_id": 1,
        "name": "Macbook Pro M1",
        "color": "Gray"
    },
    "message": "cart created",
    "status_code": 201
}
```
