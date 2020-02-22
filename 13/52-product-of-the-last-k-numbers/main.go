package main

func main() {

}

type ProductOfNumbers struct {
	prods []int
	last0 int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		last0: -1,
	}
}

func (p *ProductOfNumbers) Add(num int) {
	if num == 0 {
		p.last0 = len(p.prods)
	}
	if len(p.prods) == 0 {
		p.prods = append(p.prods, num)
	} else {
		if p.prods[len(p.prods)-1] == 0 {
			p.prods = append(p.prods, num)
		} else {
			p.prods = append(p.prods, p.prods[len(p.prods)-1]*num)
		}
	}
}

func (p *ProductOfNumbers) GetProduct(k int) int {
	start := len(p.prods) - k
	if p.last0 >= start {
		return 0
	} else if p.last0+1 == start {
		return p.prods[len(p.prods)-1]
	} else {
		first := 1
		if start > 0 {
			first = p.prods[start-1]
		}
		return p.prods[len(p.prods)-1] / first
	}
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
