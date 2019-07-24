package eazy

import "testing"

func TestLengthOfLongestSubstring(t *testing.T){
	t.Run("LengthOfLongestSubstring", func(t *testing.T) {
		length := lengthOfLongestSubstring("tmmzuxt")
		if length != 5 {
			t.Fatalf("tmmzuxt LengthOfLongestSubstring expected: 5, actual: %d",length)
		}

		length = lengthOfLongestSubstring("bbbbb")
		if length != 1 {
			t.Fatalf("tmmzuxt LengthOfLongestSubstring expected: 1, actual: %d",length)
		}

	})
}
