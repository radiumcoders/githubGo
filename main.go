package main

import "github.com/radiumcoders/githubGo/interfaces"

// "fmt"
// "os"
// "text/tabwriter"
// "github.com/radiumcoders/githubGo/structs"

func main() {
	// todo := structs.NewTodo("this is a task that is very important")
	// td := *todo
	// const padding = 4
	// w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	// fmt.Fprintf(w, "%s\t%t\t%t\t%t\t", td[0].Title, td[0].Completed, td[0].CreatedAt, td[0].UpdatedAt)
	// w.Flush()
	interfaces.InterfaceExample()
}
