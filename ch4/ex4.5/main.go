// Package ex4point5 is
// Write an in-place function to eliminate adjacent duplicatesin a []string slice.
package ex4point5

// dedupe eliminates adjacent duplicatesin a []string slice.
func dedupe(s []string) []string {
    if len(s) == 0 {
        return s
    }
    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            // Remove this s[i]
            s = append(s[:i], s[i+1:]...)
        }
    }
	return s
}
