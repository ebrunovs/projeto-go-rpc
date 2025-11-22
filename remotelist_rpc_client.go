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

	err = client.Call("RemoteList.Append", remotelist.AppendArgs{1, 10}, &ok)
	err = client.Call("RemoteList.Size", remotelist.SizeArgs{ListID: 1}, &size)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Tamanho da lista:", size)
	}

	err = client.Call("RemoteList.Append", remotelist.AppendArgs{1, 20}, &ok)
	err = client.Call("RemoteList.Size", remotelist.SizeArgs{ListID: 1}, &size)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Tamanho da lista:", size)
	}

	err = client.Call("RemoteList.Append", remotelist.AppendArgs{1, 30}, &ok)
	err = client.Call("RemoteList.Size", remotelist.SizeArgs{ListID: 1}, &size)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Tamanho da lista:", size)
	}

	err = client.Call("RemoteList.Append", remotelist.AppendArgs{2, 40}, &ok)
	err = client.Call("RemoteList.Size", remotelist.SizeArgs{ListID: 2}, &size)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Tamanho da lista:", size)
	}

	err = client.Call("RemoteList.Append", remotelist.AppendArgs{3, 50}, &ok)
	err = client.Call("RemoteList.Size", remotelist.SizeArgs{ListID: 3}, &size)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Tamanho da lista:", size)
	}

	err = client.Call("RemoteList.Remove", remotelist.RemoveArgs{1}, &removed)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento retirado:", removed)
	}
	err = client.Call("RemoteList.Size", remotelist.SizeArgs{ListID: 1}, &size)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Tamanho da lista:", size)
	}

	err = client.Call("RemoteList.Remove", remotelist.RemoveArgs{3}, &removed)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento retirado:", removed)
	}
	err = client.Call("RemoteList.Remove", remotelist.RemoveArgs{3}, &removed)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento retirado:", removed)
	}

	err = client.Call("RemoteList.Get", remotelist.GetArgs{ListID: 1, Index: 1,}, &value)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Valor encontrado: ", value)
	}
	err = client.Call("RemoteList.Get", remotelist.GetArgs{ListID: 3, Index: 2,}, &value)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Valor encontrado: ", value)
	}
}
