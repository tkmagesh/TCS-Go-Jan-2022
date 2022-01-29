// main.go

package main

func main() {
	a := App{}
	a.Initialize(
		"root",
		"rootuser",
		"public",
	)

	a.Run(":8010")
}
