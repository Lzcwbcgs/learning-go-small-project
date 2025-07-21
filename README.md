# learning-go-small-project
Small Go project for practicing basic syntax

The following business requirements description was generated with AI assistance.
# CLI-Based Small Inventory Management System (CLI-Inventory-System)  

## 1. Project Overview  

You will develop a command-line terminal application that helps small shop owners track their product inventory. The program will allow users to:  
- Add new products  
- View all products  
- Update stock quantities  
- Delete products  

All data will be stored in memory during runtime.  

**Core Value**: Helps users digitally manage product inventory through a simple tool, replacing inefficient paper-based records.  

## 2. Core Requirements  

As a shop owner, I want this tool to:  
- **Add new products**: Record product details (ID, name, price, quantity) when new items arrive.  
- **View inventory list**: See a clear list of all products for easy stock-taking.  
- **Search for a product**: Quickly find a product's details by its ID when customers inquire.  
- **Update stock**: Adjust quantities after sales (outbound) or restocking (inbound).  
- **Remove products**: Delete discontinued items from the system.  
- **Exit safely**: Close the program when finished.  

## 3. Technical Choices & Constraints  

- **Language**: Go  
- **Environment**: Command-Line Interface (CLI)  
- **Features Used**:  
  - **Variables**: Store user input and product information.  
  - **Loops**: Keep the main menu running and iterate through product lists.  
  - **Conditionals**: Execute different functions based on user input (e.g., `if/else`, `switch`).  
  - **Functions**: Encapsulate each feature (e.g., add, delete) for clean code structure.  
  - **Structs**: Define the `Product` business model.  
  - **Methods**: Attach behaviors to `Product`, such as printing its details.  
  - **Slices**: Display sorted product lists.  
  - **Map**: Store inventory with product IDs as keys for fast lookup.  
- **Restrictions**:  
  - No interfaces  
  - No concurrency (goroutines, channels)  
  - No external databases (MySQL, Redis)  
  - No complex third-party libraries  

## 4. Data Structure Design  

A core `Product` struct represents inventory items:  

```go
type Product struct {
    ID       string  // Unique product ID (e.g., "SKU001")  
    Name     string  // Product name  
    Price    float64 // Unit price  
    Quantity int     // Stock count  
}
```

**Inventory Storage**:  
- Use a `map` for fast ID-based lookups:  
  ```go
  var inventory map[string]Product  
  ```  
  - **Key**: `string` (product ID)  
  - **Value**: `Product` struct  

## 5. Functional Breakdown  

### (1) `main()`  
- Program entry point.  
- Initialize `inventory` map.  
- Use an infinite loop (`for {}`) to display the menu and process user input.  
- Call corresponding functions based on user choices (e.g., "1", "2", "q").  

### (2) `showMenu()`  
- Display the operation menu:  
  ```  
  --- Inventory Management System ---  
  1. Add Product  
  2. List All Products  
  3. Update Stock  
  4. Delete Product  
  5. Search Product  
  q. Quit  
  Enter your choice:  
  ```  

### (3) `addProduct()`  
- Guide users to input product details (ID, Name, Price, Quantity).  
- **Logic**:  
  - Check if the ID already exists in `inventory`. If yes, prompt "ID already exists."  
  - If not, create a new `Product` and add it to `inventory`.  
  - Confirm "Product added successfully."  

### (4) `listProducts()`  
- Print all products in a formatted table.  
- **Logic**:  
  - If inventory is empty, show "Inventory is empty."  
  - Otherwise, display a header (e.g., `ID | Name | Price | Quantity`) and iterate through the map.  

### (5) `updateStock()`  
- Modify a product’s quantity.  
- **Logic**:  
  - Prompt for product ID and quantity change (positive for restock, negative for sales).  
  - Verify the ID exists. If not, show "Product not found."  
  - Prevent negative stock (optional: reject if resulting quantity < 0).  
  - Confirm "Stock updated successfully."  

### (6) `deleteProduct()`  
- Remove a product by ID.  
- **Logic**:  
  - Prompt for ID, check existence, then use `delete()` to remove it from `inventory`.  
  - Show "Product deleted" or "ID not found."  

### (7) `findProduct()`  
- Search and display a single product.  
- **Logic**:  
  - Prompt for ID, search `inventory`, and print details if found.  
  - If missing, show "Product not found."  

## 6. Advanced Challenges (Post-Core Features)  

After completing the basics, try:  
- **Input Validation**: Ensure "Price" and "Quantity" are numeric (hint: `strconv`).  
- **Data Persistence**: Load/save `inventory` from/to a file (e.g., `inventory.json`) using `encoding/json`.  
- **Output Formatting**: Improve table alignment with `fmt.Printf`.  

This project provides practical Go syntax practice while building a useful tool. Happy coding! 🚀
