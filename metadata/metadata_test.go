package metadata

import (
	"context"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		mds []map[string][]string
	}
	tests := []struct {
		name string
		args args
		want Metadata
	}{
		{
			name: "hello",
			args: args{[]map[string][]string{{"hello": {"cnf"}}, {"hello2": {"go-xkit"}}}},
			want: Metadata{"hello": {"cnf"}, "hello2": {"go-xkit"}},
		},
		{
			name: "hi",
			args: args{[]map[string][]string{{"hi": {"xkit"}}, {"hi2": {"go-xkit"}}}},
			want: Metadata{"hi": {"xkit"}, "hi2": {"go-xkit"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.mds...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetadata_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		m    Metadata
		args args
		want string
	}{
		{
			name: "xkit",
			m:    Metadata{"xkit": {"value"}, "env": {"dev"}},
			args: args{key: "xkit"},
			want: "value",
		},
		{
			name: "env",
			m:    Metadata{"xkit": {"value"}, "env": {"dev"}},
			args: args{key: "env"},
			want: "dev",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Get(tt.args.key); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetadata_Values(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		m    Metadata
		args args
		want []string
	}{
		{
			name: "xkit",
			m:    Metadata{"xkit": {"value", "value2"}, "env": {"dev"}},
			args: args{key: "xkit"},
			want: []string{"value", "value2"},
		},
		{
			name: "env",
			m:    Metadata{"xkit": {"value", "value2"}, "env": {"dev"}},
			args: args{key: "env"},
			want: []string{"dev"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Values(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetadata_Set(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		m    Metadata
		args args
		want Metadata
	}{
		{
			name: "xkit",
			m:    Metadata{},
			args: args{key: "hello", value: "xkit"},
			want: Metadata{"hello": {"xkit"}},
		},
		{
			name: "env",
			m:    Metadata{"hello": {"xkit"}},
			args: args{key: "env", value: "pro"},
			want: Metadata{"hello": {"xkit"}, "env": {"pro"}},
		},
		{
			name: "empty",
			m:    Metadata{},
			args: args{key: "", value: ""},
			want: Metadata{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Set(tt.args.key, tt.args.value)
			if !reflect.DeepEqual(tt.m, tt.want) {
				t.Errorf("Set() = %v, want %v", tt.m, tt.want)
			}
		})
	}
}

func TestMetadata_Add(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		m    Metadata
		args args
		want Metadata
	}{
		{
			name: "xkit",
			m:    Metadata{},
			args: args{key: "hello", value: "xkit"},
			want: Metadata{"hello": {"xkit"}},
		},
		{
			name: "env",
			m:    Metadata{"hello": {"xkit"}},
			args: args{key: "hello", value: "again"},
			want: Metadata{"hello": {"xkit", "again"}},
		},
		{
			name: "empty",
			m:    Metadata{},
			args: args{key: "", value: ""},
			want: Metadata{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Add(tt.args.key, tt.args.value)
			if !reflect.DeepEqual(tt.m, tt.want) {
				t.Errorf("Set() = %v, want %v", tt.m, tt.want)
			}
		})
	}
}

func TestClientContext(t *testing.T) {
	type args struct {
		ctx context.Context
		md  Metadata
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "xkit",
			args: args{context.Background(), Metadata{"hello": {"xkit"}, "xkit": {"https://go-xkit.dev"}}},
		},
		{
			name: "hello",
			args: args{context.Background(), Metadata{"hello": {"xkit"}, "hello2": {"https://go-xkit.dev"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := NewClientContext(tt.args.ctx, tt.args.md)
			m, ok := FromClientContext(ctx)
			if !ok {
				t.Errorf("FromClientContext() = %v, want %v", ok, true)
			}

			if !reflect.DeepEqual(m, tt.args.md) {
				t.Errorf("meta = %v, want %v", m, tt.args.md)
			}
		})
	}
}

func TestServerContext(t *testing.T) {
	type args struct {
		ctx context.Context
		md  Metadata
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "xkit",
			args: args{context.Background(), Metadata{"hello": {"xkit"}, "xkit": {"https://go-xkit.dev"}}},
		},
		{
			name: "hello",
			args: args{context.Background(), Metadata{"hello": {"xkit"}, "hello2": {"https://go-xkit.dev"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := NewServerContext(tt.args.ctx, tt.args.md)
			m, ok := FromServerContext(ctx)
			if !ok {
				t.Errorf("FromServerContext() = %v, want %v", ok, true)
			}

			if !reflect.DeepEqual(m, tt.args.md) {
				t.Errorf("meta = %v, want %v", m, tt.args.md)
			}
		})
	}
}

func TestAppendToClientContext(t *testing.T) {
	type args struct {
		md Metadata
		kv []string
	}
	tests := []struct {
		name string
		args args
		want Metadata
	}{
		{
			name: "xkit",
			args: args{Metadata{}, []string{"hello", "xkit", "env", "dev"}},
			want: Metadata{"hello": {"xkit"}, "env": {"dev"}},
		},
		{
			name: "hello",
			args: args{Metadata{"hi": {"https://go-xkit.dev/"}}, []string{"hello", "xkit", "env", "dev"}},
			want: Metadata{"hello": {"xkit"}, "env": {"dev"}, "hi": {"https://go-xkit.dev/"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := NewClientContext(context.Background(), tt.args.md)
			ctx = AppendToClientContext(ctx, tt.args.kv...)
			md, ok := FromClientContext(ctx)
			if !ok {
				t.Errorf("FromServerContext() = %v, want %v", ok, true)
			}
			if !reflect.DeepEqual(md, tt.want) {
				t.Errorf("metadata = %v, want %v", md, tt.want)
			}
		})
	}
}

// nolint directives: sa5012
func TestAppendToClientContextThatPanics(t *testing.T) {
	kvs := []string{"hello", "xkit", "env"}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("append to client context singular kvs did not panic")
		}
	}()
	ctx := NewClientContext(context.Background(), Metadata{})
	ctx = AppendToClientContext(ctx, kvs...)
	md, ok := FromClientContext(ctx)
	if !ok {
		t.Errorf("FromServerContext() = %v, want %v", ok, true)
	}
	if !reflect.DeepEqual(md, Metadata{}) {
		t.Errorf("metadata = %v, want %v", md, Metadata{})
	}
}

func TestMergeToClientContext(t *testing.T) {
	type args struct {
		md       Metadata
		appendMd Metadata
	}
	tests := []struct {
		name string
		args args
		want Metadata
	}{
		{
			name: "xkit",
			args: args{Metadata{}, Metadata{"hello": {"xkit"}, "env": {"dev"}}},
			want: Metadata{"hello": {"xkit"}, "env": {"dev"}},
		},
		{
			name: "hello",
			args: args{Metadata{"hi": {"https://go-xkit.dev/"}}, Metadata{"hello": {"xkit"}, "env": {"dev"}}},
			want: Metadata{"hello": {"xkit"}, "env": {"dev"}, "hi": {"https://go-xkit.dev/"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := NewClientContext(context.Background(), tt.args.md)
			ctx = MergeToClientContext(ctx, tt.args.appendMd)
			md, ok := FromClientContext(ctx)
			if !ok {
				t.Errorf("FromServerContext() = %v, want %v", ok, true)
			}
			if !reflect.DeepEqual(md, tt.want) {
				t.Errorf("metadata = %v, want %v", md, tt.want)
			}
		})
	}
}

func TestMetadata_Range(t *testing.T) {
	md := Metadata{"xkit": {"xkit"}, "https://go-xkit.dev/": {"https://go-xkit.dev/"}, "go-xkit": {"go-xkit"}}
	tmp := Metadata{}
	md.Range(func(k string, v []string) bool {
		if k == "https://go-xkit.dev/" || k == "xkit" {
			tmp[k] = v
		}
		return true
	})
	if !reflect.DeepEqual(tmp, Metadata{"https://go-xkit.dev/": {"https://go-xkit.dev/"}, "xkit": {"xkit"}}) {
		t.Errorf("metadata = %v, want %v", tmp, Metadata{"https://go-xkit.dev/": {"https://go-xkit.dev/"}, "xkit": {"xkit"}})
	}
	tmp = Metadata{}
	md.Range(func(string, []string) bool {
		return false
	})
	if !reflect.DeepEqual(tmp, Metadata{}) {
		t.Errorf("metadata = %v, want %v", tmp, Metadata{})
	}
}

func TestMetadata_Clone(t *testing.T) {
	tests := []struct {
		name string
		m    Metadata
		want Metadata
	}{
		{
			name: "xkit",
			m:    Metadata{"xkit": {"xkit"}, "https://go-xkit.dev/": {"https://go-xkit.dev/"}, "go-xkit": {"go-xkit"}},
			want: Metadata{"xkit": {"xkit"}, "https://go-xkit.dev/": {"https://go-xkit.dev/"}, "go-xkit": {"go-xkit"}},
		},
		{
			name: "go",
			m:    Metadata{"language": {"golang"}},
			want: Metadata{"language": {"golang"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.Clone()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
			got["xkit"] = []string{"go"}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("want got != want got %v want %v", got, tt.want)
			}
		})
	}
}
