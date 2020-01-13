package main

func main() {

}

type Logger struct {
	last map[string]int
}

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{last: map[string]int{}}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	v, ok := l.last[message]
	if !ok || v <= timestamp-10 {
		l.last[message] = timestamp
		return true
	}
	return false
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
