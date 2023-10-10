
# Tasty Bites


## Background

Made as a personal project for Non-Graded Challenge

## Installation

Make sure you have initialized ```go mod init``` on the same folder of your project.
Then, copy the following command:
```
go get github.com/DaniMunandar/tastybites

```
## Usage
Use ```tastybites.Start()``` to use the module.

``` 
import "github.com/DaniMunandar/tastybites"

func main() {
	tastybites.Start()
}
```

You will be prompted to enter a database connection (currently for private usage)

``` 
Enter connection with the following format:
[USERNAME]:[PASSWORD]@tcp([HOST]:[PORT])/[DBNAME]
```

Once the connection successfully establised, the CLI app is ready to use!

```
Options:
1. Add Employee
2. View Menu Items
3. Process Order
4. Exit
Enter your option: 
```


