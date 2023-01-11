package presenter

import (
	"context"
	"testing"

	"github.com/joaquinicolas/iris-bot/src/domain/entities"
	"github.com/stretchr/testify/assert"
)

func Test_botPresenter_RespondWithProducts(t *testing.T) {
	const regexPattern = `\d\.\s*[a-zA-Z]+:.+`
	type fields struct {
		ctx context.Context
	}
	type args struct {
		products []entities.Product
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		count  int
	}{
		{
			name:   "respond with products",
			fields: fields{ctx: context.Background()},
			args:   args{products: []entities.Product{{Name: "test", Price: "30$"}}},
			want:   regexPattern,
			count:  1,
		},
		{
			name:   "respond with empty string",
			fields: fields{ctx: context.Background()},
			args:   args{products: []entities.Product{}},
			want:   "",
			count:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &botPresenter{
				ctx: tt.fields.ctx,
			}
			got, c := b.RespondWithProducts(tt.args.products)
			if c != tt.count {
				t.Errorf("botPresenter.RespondWithProducts() expected %d products, got %d", tt.count, c)
				return
			}
			if c > 0 {
				assert.Regexp(t, regexPattern, got)
			} else {
				assert.Equal(t, "", got)
			}
		})
	}
}
