package main

import (
	"fmt"
	"net/rpc"
	"ifpb/remotelist/pkg"
)

func main() {
	client, err := rpc.Dial("tcp", ":5000")
	if err != nil {
		fmt.Print("dialing:", err)
	}

	// Synchronous call
	var ok bool

	err = client.Call("RemoteList.Append", remotelist.AppendArgs{"A", 10}, &ok)
	err = client.Call("RemoteList.Append", remotelist.AppendArgs{"A", 20}, &ok)
	err = client.Call("RemoteList.Append", remotelist.AppendArgs{"A", 30}, &ok)
	err = client.Call("RemoteList.Append", remotelist.AppendArgs{"B", 40}, &ok)
	err = client.Call("RemoteList.Append", remotelist.AppendArgs{"abc", 50}, &ok)

	// err = client.Call("RemoteList.Remove", 0, &reply_i)
	// if err != nil {
	// 	fmt.Print("Error:", err)
	// } else {
	// 	fmt.Println("Elemento retirado:", reply_i)
	// }
	// err = client.Call("RemoteList.Remove", 0, &reply_i)
	// if err != nil {
	// 	fmt.Print("Error:", err)
	// } else {
	// 	fmt.Println("Elemento retirado:", reply_i)
	// }
}
