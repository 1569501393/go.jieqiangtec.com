package main

import (
	"reflect"
	"testing"
)

func TestCountInto(t *testing.T) {
	type args struct {
		a     Appender
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestList_Append(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		l    List
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestList_Len(t *testing.T) {
	tests := []struct {
		name string
		l    List
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongEnough(t *testing.T) {
	type args struct {
		l Lener
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongEnough(tt.args.l); got != tt.want {
				t.Errorf("LongEnough() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNode(t *testing.T) {
	type args struct {
		left  *Node
		right *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_SetData(t *testing.T) {
	type fields struct {
		le   *Node
		data interface{}
		ri   *Node
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				le:   tt.fields.le,
				data: tt.fields.data,
				ri:   tt.fields.ri,
			}
		})
	}
}

func TestRectangle_Area(t *testing.T) {
	type fields struct {
		length float32
		width  float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				length: tt.fields.length,
				width:  tt.fields.width,
			}
			if got := r.Area(); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare_Area(t *testing.T) {
	type fields struct {
		side float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sq := &Square{
				side: tt.fields.side,
			}
			if got := sq.Area(); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_car_getValue(t *testing.T) {
	type fields struct {
		make  string
		model string
		price float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := car{
				make:  tt.fields.make,
				model: tt.fields.model,
				price: tt.fields.price,
			}
			if got := c.getValue(); got != tt.want {
				t.Errorf("getValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_showValue(t *testing.T) {
	type args struct {
		asset valuable
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_stockPosition_getValue(t *testing.T) {
	type fields struct {
		ticker     string
		sharePrice float32
		count      float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stockPosition{
				ticker:     tt.fields.ticker,
				sharePrice: tt.fields.sharePrice,
				count:      tt.fields.count,
			}
			if got := s.getValue(); got != tt.want {
				t.Errorf("getValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
