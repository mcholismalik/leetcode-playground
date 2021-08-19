package naryTreeLevelOrderTraversal

func Seed1() *Node {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 2,
				Children: []*Node{
					{
						Val:      5,
						Children: nil,
					},
					{
						Val:      6,
						Children: nil,
					},
				},
			},
			{
				Val: 3,
				Children: []*Node{
					{
						Val:      7,
						Children: nil,
					},
					{
						Val: 8,
						Children: []*Node{
							{
								Val:      9,
								Children: nil,
							},
							{
								Val:      10,
								Children: nil,
							},
						},
					},
				},
			},
			{
				Val: 4,
				Children: []*Node{
					{
						Val: 11,
						Children: []*Node{
							{
								Val: 12,
								Children: []*Node{
									{
										Val:      13,
										Children: nil,
									},
								},
							},
						},
					},
				},
			},
		}}
	return root
}

func Seed2() *Node {
	root := &Node{
		Val:      1,
		Children: nil,
	}
	return root
}

func Seed3() *Node {
	return nil
}
