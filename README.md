## Neo

A basic web server which performs the following operations on matrices.
Given an uploaded csv file

```
1,2,3
4,5,6
7,8,9
```

1. Echo
   - Return the matrix as a string in matrix format.
   ```
   1,2,3
   4,5,6
   7,8,9
   ```
2. Invert
   - Return the matrix as a string in matrix format where the columns and rows are inverted
   ```
   // Expected output
   1,4,7
   2,5,8
   3,6,9
   ```
3. Flatten
   - Return the matrix as a 1 line string, with values separated by commas.
   ```
   1,2,3,4,5,6,7,8,9
   ```
4. Sum
   - Return the sum of the integers in the matrix
   ```
   45
   ```
5. Multiply
   - Return the product of the integers in the matrix
   ```
   362880
   ```

- The input file to these functions is a square matrix of any dimension, i.e. the number of rows is equal to the number of columns
- Each value is an integer, and there is no header row

## To Run:

- To run the server

```
go run .
```

- Send request

```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
curl -F 'file=@/path/matrix.csv' "localhost:8080/invert"
curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"
curl -F 'file=@/path/matrix.csv' "localhost:8080/sum"
curl -F 'file=@/path/matrix.csv' "localhost:8080/multiply"
```

- To run tests

```
go test ./handlers -v
```
