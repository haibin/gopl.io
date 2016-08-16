// Package ex4point4 is
// write a version of rotate that operates in a single pass.
package ex4point4

// rotate rotates s left by n positions.
func rotate(s []int, n int) {
	for j := 0; j < n; j++ {
		tmp := s[0]
		for i := 0; i < len(s)-1; i++ {
			s[i] = s[i+1]
		}
		s[len(s)-1] = tmp
	}
}
