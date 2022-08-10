package sort

import (
	"reflect"
	"testing"
)

func Test_maxHeap_Insert(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test-case-1",
			args: args{
				val: 2,
			},
		},
		{
			name: "test-case-2",
			args: args{
				val: 3,
			},
		},
		{
			name: "test-case-3",
			args: args{
				val: 4,
			},
		},
		{
			name: "test-case-4",
			args: args{
				val: 52,
			},
		},
	}
	heap := NewMaxHeap(4)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap.Insert(tt.args.val)
			t.Logf("tt name %s heap is %v", tt.name, heap)
		})
	}
	t.Logf("result heap is %v", heap)
}

func Test_maxHeap_ExtraMax(t *testing.T) {
	type fields struct {
		Data  []int
		Count int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "test-case-1",
			fields: fields{
				Data:  []int{23, 8, 21, 5, 6, 9, 11},
				Count: 7,
			},
			want: 23,
		},
		{
			name: "test-case-2",
			fields: fields{
				Data:  []int{21, 8, 11, 5, 6, 9},
				Count: 6,
			},
			want: 21,
		},
		{
			name: "test-case-3",
			fields: fields{
				Data:  []int{21, 3},
				Count: 2,
			},
			want: 21,
		},
		{
			name: "test-case-4",
			fields: fields{
				Data:  []int{3},
				Count: 1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := &maxHeap{
				Data:  tt.fields.Data,
				Count: tt.fields.Count,
			}
			if got := heap.ExtraMax(); got != tt.want {
				t.Errorf("maxHeap.ExtraMax() = %v, want %v", got, tt.want)
			} else {
				t.Logf("got %d", got)
			}
			t.Logf("result heap is %v", heap)
		})
	}
}

func TestMaxHeapipy(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test-case-1",
			args: args{
				arr: []int{15, 17, 19, 13, 22, 16, 28, 30, 41, 62},
			},
			want: []int{62, 41, 28, 30, 22, 16, 19, 15, 13, 17},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxHeapipy(tt.args.arr).Data; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxHeapipy() = %v, want %v", got, tt.want)
			}
		})
	}
}
