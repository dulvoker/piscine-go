package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var lis *NodeL
	lis = l.Head
	for lis != nil {
		if CompStr(lis.Next.Data, data_ref) {
			if lis.Next.Next != nil {
				lis.Next = lis.Next.Next
			} else {
				lis.Next = nil
			}
		}
		lis = lis.Next
	}
}
