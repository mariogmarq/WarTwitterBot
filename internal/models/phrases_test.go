package models

import (
	"testing"
)

func TestPhrase_MapToPlayers(t *testing.T) {
	type fields struct {
		Phrase string
		N      int
	}
	type args struct {
		fighters []FighterApi
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should rise error in case does not match",
			fields: fields{
				Phrase: "$1 kills $2",
				N:      2,
			},
			args: args{
				fighters: []FighterApi{
					{Name: "Mario"},
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should replace properly",
			fields: fields{
				Phrase: "$1 kills $2",
				N:      2,
			},
			args: args{
				fighters: []FighterApi{
					{Name: "Mario"}, {Name: "Dani"},
				},
			},
			want:    "Mario kills Dani",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Phrase{
				Phrase: tt.fields.Phrase,
				N:      tt.fields.N,
			}
			got, err := p.MapToPlayers(tt.args.fighters...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Phrase.MapToPlayers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Phrase.MapToPlayers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPhrase(t *testing.T) {
	type args struct {
		phrase string
	}
	tests := []struct {
		name    string
		args    args
		want    int //Only check N
		wantErr bool
	}{
		{
			name: "Works fine in a proper phrase",
			args: args{
				phrase: "$1 kills $2",
			},
			want:    2,
			wantErr: false,
		}, {
			name: "$X are skipped naturals",
			args: args{
				phrase: "$1 kills $3",
			},
			want:    0,
			wantErr: true,
		}, {
			name: "There is only one $1",
			args: args{
				phrase: "$1 kills $1",
			},
			want:    0,
			wantErr: true,
		}, {
			name: "Other case where it works fine",
			args: args{
				phrase: "$1 killed $3 with the help of $2",
			},
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPhrase(tt.args.phrase)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPhrase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.N != tt.want {
				t.Errorf("NewPhrase() = %v, want %v", got, tt.want)
			}
		})
	}
}
