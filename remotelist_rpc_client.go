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

	var ok bool
	var removed int
	var value int
	var size int

	err = client.Call("RemoteList.Append", remotelist.AppendArgs{"A", 10}, &ok)
	err = client.Call("RemoteList.Size", remotelist.SizeArgs{ListID: "A"}, &size)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Tamanho da lista:", size)
	}

	err = client.Call("RemoteList.Append", remotelist.AppendArgs{"A", 20}, &ok)
	err = client.Call("RemoteList.Size", remotelist.SizeArgs{ListID: "A"}, &size)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Tamanho da lista:", size)
	}

	err = client.Call("RemoteList.Append", remotelist.AppendArgs{"A", 30}, &ok)
	err = client.Call("RemoteList.Size", remotelist.SizeArgs{ListID: "A"}, &size)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Tamanho da lista:", size)
	}

	err = client.Call("RemoteList.Append", remotelist.AppendArgs{"B", 40}, &ok)
	err = client.Call("RemoteList.Size", remotelist.SizeArgs{ListID: "B"}, &size)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Tamanho da lista:", size)
	}

	err = client.Call("RemoteList.Append", remotelist.AppendArgs{"abc", 50}, &ok)
	err = client.Call("RemoteList.Size", remotelist.SizeArgs{ListID: "abc"}, &size)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Tamanho da lista:", size)
	}

	err = client.Call("RemoteList.Remove", remotelist.RemoveArgs{"A"}, &removed)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento retirado:", removed)
	}
	err = client.Call("RemoteList.Size", remotelist.SizeArgs{ListID: "A"}, &size)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Tamanho da lista:", size)
	}

	err = client.Call("RemoteList.Remove", remotelist.RemoveArgs{"abc"}, &removed)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento retirado:", removed)
	}
	err = client.Call("RemoteList.Remove", remotelist.RemoveArgs{"abc"}, &removed)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento retirado:", removed)
	}

	err = client.Call("RemoteList.Get", remotelist.GetArgs{ListID: "A", Index: 1,}, &value)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Valor encontrado: ", value)
	}
	err = client.Call("RemoteList.Get", remotelist.GetArgs{ListID: "abc", Index: 2,}, &value)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Valor encontrado: ", value)
	}
}
