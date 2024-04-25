# Polysurance - Full Stack Developer Test

This application solves the first scenario from Polysurance Developer Test

### Statement 

> During the development of a sales software product, you have been given the following task:
>
>
>
>Write a set of functions that calculate:
>
>Total sales before discount is applied;
>Total sales after discount code is applied;
>Total amount of money lost via customer using discount codes;
>Average discount per customer as a percentage;
> 
>
>The attached JSON files contain objects with information about discount codes, item prices and order details.
>
>
>
>Given the supplied data payloads, please write a solution to the above scenario in any language, keeping it  efficient and readable. Please include the outputs of the functions in your reply.

### Solution

This project aims to solve the statement above by implementing an application in Go that reads input files and builds a data structure with the total values - without discount - of each order, called `OrderInfo`. After processing the files, a slice of `OrderInfo` is obtained and to get the answers to the questions, we simply iterate over the elements of this slice, performing the necessary calculations. The logic can be found in the file `/sales/sales.go`.

### How to Run
1. Clone the repo:
```bash
 git clone https://github.com/ducho-metson/polysurance-test.git
```
2. Go to application dir
```bash
cd polysurance-test
```

3. Run the application
    a. Using Local Golang 
    ```bash
    go run main.go
    ```
    b. Using Docker
    ```bash
    docker-compose up --build
    ```

