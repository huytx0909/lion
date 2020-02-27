package tests

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

//func TestKeyOp_Generate(t *testing.T) {
//	type args struct {
//		x int
//		y int
//	}
//
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		{
//			name: "success",
//			args: args{
//				x: 5,
//				y: 50,
//			},
//			want: "5_50",
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			kp := GetKeyOperator()
//			if got := kp.Generate(tt.args.x, tt.args.y); got != tt.want {
//				t.Errorf("keyOp.Generate() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestKeyOp_Degenerate(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		key string
	}

	testCases := []struct {
		name     string
		argument args
	}{
		{
			name: "case 1",
			argument: args{
				key: "123_456",
			},
		},
		{
			name: "case 1",
			argument: args{
				key: "678_9101112",
			},
		},
	}

	generator := GetKeyOperator()

	for _, tc := range testCases {
		pre, post, err := generator.Degenerate(tc.argument.key)

		wantedSlice := strings.Split(tc.argument.key, "_")

		wantedPre := wantedSlice[0]
		wantedPost := wantedSlice[1]

		wantedPreInt, _ := strconv.Atoi(wantedPre)
		wantedPostInt, _ := strconv.Atoi(wantedPost)

		assert.Equal(wantedPreInt, pre, "They should be equal")
		assert.Equal(wantedPostInt, post, "They should be equal")
		assert.Nil(err)
	}

}
