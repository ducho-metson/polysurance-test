# Polysurance - Full Stack Developer Test

This application solves the first and second scenarios from Polysurance Developer Test

### Statement 

##### Scenario 1 

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

##### Scenario 2
> The product team has added a new discount code “WINTERMADNESS” with a value of 10%. This  new code stacks with existing discount codes additively. For example, if using both “SALE30” and “WINTERMADNESS” the discount would be 40%.
>
> Please calculate the same summary statistics of the orders, as calculated previously, using the attached json files.

### Solution

This project aims to solve the statements above by implementing an application in Go that reads input files and builds a data structure with the total values - without discount - and applied discount from each order, called `OrderInfo`. After processing the files, a slice of `OrderInfo` is obtained and to get the answers to the questions, we simply iterate over the elements of this slice, performing the necessary calculations. The logic can be found in the file `/sales/sales.go`.

### How to Run
1. Clone the repo:
```bash
 git clone https://github.com/ducho-metson/polysurance-test.git
```
2. Go to application dir
```bash
cd polysurance-test
```

3. Run the application passing the desired scenario 
    a. Using Local Golang 
    ```bash
    go run main.go part1 # for scenario 1
    go run main.go part2 # for scenario 2
    ```
    b. Using Docker
    ```bash
    set ARGUMENT=part1 && docker-compose up --build # for scenario 1
    set ARGUMENT=part2 && docker-compose up --build # for scenario 2
    ```

