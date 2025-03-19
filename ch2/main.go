package main

import "fmt"

// import "fmt"

// // var (
// // 	x    int
// // 	y        = 20
// // 	z    int = 30
// // 	d, e     = 40, "hello"
// // 	f, g string
// // )

// const x int64 = 10
// const (
// 	idKey = "id"
// 	nameKey = "name"
// )

// const z = 20 * 10

// func main() {
// 	// Print the values of the variables
// 	// fmt.Println(x, y, z, d, e, f, g)

// 	// // Assign values to f and g
// 	// f = "world"
// 	// g = "!"

// 	// // Print the updated values of f and g
// 	// fmt.Println(f, g)

// 	// // // Print the types of the variables
// 	// // println("Type of x:", typeof(x))
// 	// // println("Type of y:", typeof(y))
// 	// // println("Type of z:", typeof(z))
// 	// // println("Type of d:", typeof(d))
// 	// // println("Type of e:", typeof(e))
// 	// // println("Type of f:", typeof(f))
// 	// // println("Type of g:", typeof(g))

// 	// x := 10

// 	// fmt.Println(x)
// 	// x, y := 30, "hello"

// 	// // Print the values of the new variables
// 	// fmt.Println(x, y)

// 	const y = "hello"

// 	fmt.Println(x)
// 	fmt.Println(y)

// 	x = x + 1
// 	y = "bye"

// }

type Config struct {
	appName string
}

func NewConfig() *Config {
	return &Config{
		appName: "myapp"}
}

func (c *Config) GetAppName() string {
	return c.appName
}

func main() {
	cfg := NewConfig()
	fmt.Println("App Name:", cfg.GetAppName())
}