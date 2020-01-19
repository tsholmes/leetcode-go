package main

func main() {

}

type PhoneDirectory struct {
	used []bool
}

/** Initialize your data structure here
  @param maxNumbers - The maximum numbers that can be stored in the phone directory. */
func Constructor(maxNumbers int) PhoneDirectory {
	return PhoneDirectory{used: make([]bool, maxNumbers)}
}

/** Provide a number which is not assigned to anyone.
  @return - Return an available number. Return -1 if none is available. */
func (p *PhoneDirectory) Get() int {
	for i, used := range p.used {
		if !used {
			p.used[i] = true
			return i
		}
	}
	return -1
}

/** Check if a number is available or not. */
func (p *PhoneDirectory) Check(number int) bool {
	return !p.used[number]
}

/** Recycle or release a number. */
func (p *PhoneDirectory) Release(number int) {
	p.used[number] = false
}

/**
 * Your PhoneDirectory object will be instantiated and called as such:
 * obj := Constructor(maxNumbers);
 * param_1 := obj.Get();
 * param_2 := obj.Check(number);
 * obj.Release(number);
 */
