package sort
type Node struct{
	Data int
	Next *Node
}

/***************
* 返回第一个节点
* h 头结点
 ***************/
func GetFirst(h *Node) *Node{
	if h.Next == nil{
		return nil
	}
	return h.Next
}

/***************
* 返回最后一个节点
 */
func GetLast(h *Node) *Node{
	if h.Next == nil{
		return nil
	}
	var i *Node = h
	for i.Next != nil{
		i = i.Next
		if i.Next == nil{
			return i
		}
	}
	return nil
}

/***************
* 返回长度
 ***************/
func Length(h *Node) int{
	var i int = 0
	n := h
	for n.Next != nil{
		i++
		n = n.Next
	}
	return i
}

/***************
* 插入节点
 ***************/
func Insert(h *Node, d *Node, i int) *Node{
	var node *Node = h
	var j int = 1
	for node.Next != nil{
		if i == j{
			d.Next = node.Next
			node.Next = d
		}
		node = node.Next
		j++
	}
	return h
}
/***************
* 删除节点
 ***************/
func DeleteNode(h *Node, i int) *Node{
	var node *Node = h
	var j int = 1
	for node.Next != nil{
		if i == j{
			node.Next = node.Next.Next
			break
		}
		node = node.Next
		j++
	}
	return h
}
/***************
* 获取节点
 ***************/
func Get(h *Node, i int) *Node{
	var node *Node = h
	var j int = 1
	for node.Next != nil{
		if i == j{
			return node
		}
		node = node.Next
		j++
	}
	return nil
}
