package main

func main() {
	result := add(1, 2)

}

func add(a, b interface{}) interface{} {
	aInt, aOk := a.(int)
	bInt, bOk := b.(int)

	if aOk && bOk {
		return aInt + bInt
	}

	aString, aOk := a.(string)
	bString, bOk := b.(string)

	if aOk && bOk {
		return aString + bString
	}

	return nil
}
