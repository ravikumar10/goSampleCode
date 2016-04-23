package main
import "fmt"


type Tree struct {
    Value int
    Branches map[string]Tree
}


func TreeSum( tree Tree ) int  {
	sum := tree.Value;
	for _, val := range tree.Branches {
		sum += TreeSum(val)
	}
	return sum;
}

type iInt int

func (max iInt) Iter () <-chan iInt {
	ch := make(chan iInt);
	go func () {
		m := int(max)
		for i := 0; i <= m; i++ {
			ch <- iInt(i)
		}
		close(ch)
	} ();
	return ch
}

func TreeWalker( tree Tree, f func(Tree) )  {
	f(tree)
	for _, val := range tree.Branches {
		TreeWalker(val,f)
	}
}

func TreeIter( tree Tree ) <-chan Tree {
	ch := make(chan Tree);
	go func () {
		TreeWalker( tree, func(t Tree) {
			ch <- t
		})
		close(ch)
	} ();
	return ch
}

func (tree Tree) Iter() <-chan Tree {
	return TreeIter( tree )
}
func main(){
  stump := map[string]Tree{}

	tree := Tree{0,
		map[string]Tree{
			"a":Tree{1,stump},
			"b":Tree{2, map[string]Tree{
					"h":Tree{8,stump},
					"i":Tree{9,stump},
					"j":Tree{10,stump},
				},
			},
			"c":Tree{3,stump},
			"d":Tree{4, map[string]Tree{
					"e":Tree{5,stump},
					"f":Tree{6,stump},
					"g":Tree{7,stump},
				},
			},
		},
	}

	fmt.Printf("Treesum %d\n", TreeSum( tree ))

	for i := range iInt(6).Iter() {
		fmt.Printf("Wow! %v\n", i)
	}

	// Call Back walker
	TreeWalker(tree, func(t Tree) { fmt.Printf("Node value %d\n",t.Value) } )

	for tree := range tree.Iter() {
		fmt.Printf("Now via Iter Node value %d\n",tree.Value)
	}
}
