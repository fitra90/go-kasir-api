# TestGo API Documentation

This project provides a simple REST API for managing Products and Categories using Go and PostgreSQL (Supabase).

## Base URL
`http://localhost:8080`

## Endpoints

### Health Check
Check if the API is running.

- **URL:** `/health`
- **Method:** `GET`
- **Response:**
  ```json
  {
      "status": "OK",
      "message": "API Running"
  }
  ```

---

### Categories

#### 1. Get All Categories
Retrieve a list of all categories.

- **URL:** `/api/category`
- **Method:** `GET`
- **Response:**
  ```json
  [
      {
          "id": 1,
          "name": "Beverages",
          "description": "Soft drinks and juices"
      },
      {
          "id": 2,
          "name": "Snacks",
          "description": "Chips and crackers"
      }
  ]
  ```

#### 2. Get Category by ID
Retrieve details of a specific category.

- **URL:** `/api/category/{id}`
- **Method:** `GET`
- **Response:**
  ```json
  {
      "id": 1,
      "name": "Beverages",
      "description": "Soft drinks and juices"
  }
  ```

#### 3. Create Category
Create a new category.

- **URL:** `/api/category`
- **Method:** `POST`
- **Body:**
  ```json
  {
      "name": "Electronics",
      "description": "Gadgets and devices"
  }
  ```
- **Response:**
  ```json
  {
      "id": 3,
      "name": "Electronics",
      "description": "Gadgets and devices"
  }
  ```

#### 4. Update Category
Update an existing category.

- **URL:** `/api/category/{id}`
- **Method:** `PUT`
- **Body:**
  ```json
  {
      "name": "Home Appliances",
      "description": "Furniture and decor"
  }
  ```
- **Response:**
  ```json
  {
      "id": 3,
      "name": "Home Appliances",
      "description": "Furniture and decor"
  }
  ```

#### 5. Delete Category
Delete a category.

- **URL:** `/api/category/{id}`
- **Method:** `DELETE`
- **Response:** (204 No Content or success message depending on implementation)

---

### Products

#### 1. Get All Products
Retrieve a list of all products, including their category details.

- **URL:** `/api/produk`
- **Method:** `GET`
- **Response:**
  ```json
  [
      {
          "id": 1,
          "name": "Sprite",
          "price": 4500,
          "stock": 10,
          "category_id": 1,
          "category": {
              "id": 1,
              "name": "Beverages",
              "description": "Soft drinks and juices"
          }
      }
  ]
  ```

#### 2. Get Product by ID
Retrieve details of a specific product.

- **URL:** `/api/produk/{id}`
- **Method:** `GET`
- **Response:**
  ```json
  {
      "id": 1,
      "name": "Sprite",
      "price": 4500,
      "stock": 10,
      "category_id": 1,
      "category": {
          "id": 1,
          "name": "Beverages",
          "description": "Soft drinks and juices"
      }
  }
  ```

#### 3. Create Product
Create a new product. Note that the response format mirrors the input data.

- **URL:** `/api/produk`
- **Method:** `POST`
- **Body:**
  ```json
  {
      "name": "Coca Cola",
      "price": 5000,
      "stock": 20,
      "category_id": 1
  }
  ```
- **Response:**
  ```json
  {
      "name": "Coca Cola",
      "price": 5000,
      "stock": 20,
      "category_id": 1
  }
  ```

#### 4. Update Product
Update an existing product.

- **URL:** `/api/produk/{id}`
- **Method:** `PUT`
- **Body:**
  ```json
  {
      "name": "Coca Cola Zero",
      "price": 5500,
      "stock": 15,
      "category_id": 1
  }
  ```
- **Response:**
  ```json
  {
      "id": 2,
      "name": "Coca Cola Zero",
      "price": 5500,
      "stock": 15,
      "category_id": 1,
      "category": { ... } 
  }
  ```
  *(Note: Response structure for update may vary slightly based on implementation details)*

#### 5. Delete Product
Delete a product.

- **URL:** `/api/produk/{id}`
- **Method:** `DELETE`
- **Response:** (204 No Content or success message)
