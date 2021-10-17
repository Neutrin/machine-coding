package user

import "testing"

func Test_GetMark(t *testing.T) {
	tests := []struct {
		name     string
		object   *User
		expected rune
	}{
		{
			name: "mark x",
			object: &User{
				name:     "nitin",
				password: "nitin",
				mark:     rune('X'),
			},
			expected: rune('X'),
		},
		{
			name: "mark O",
			object: &User{
				name:     "neutrin",
				password: "neutrin",
				mark:     rune('O'),
			},
			expected: rune('O'),
		},
	}
	for _, curTest := range tests {
		t.Run(curTest.name, func(t *testing.T) {
			if got := curTest.object.GetMark(); got != curTest.expected {
				t.Errorf("got = %v, wanted = %v", got, curTest.expected)
			}
		})
	}
}
