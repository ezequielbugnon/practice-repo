package bitmanipulation

/*
Binary representation of a given number
Last Updated : 17 Mar, 2025
Given an integer n, the task is to print the binary representation of the number.

Note: The given number will be maximum of 32 bits, so append 0's to the left if the result string is smaller than 30 length.

Examples:

Input: n = 2
Output: 00000000000000000000000000000010

Input: n = 0
Output: 00000000000000000000000000000000
*/

func BinaryRepresentation(n int) string {
	var result = ""

	for i := 32; i >= 0; i-- {
		if (n & (1 << i)) != 0 {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}
