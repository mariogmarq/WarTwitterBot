package utils

import "testing"

func TestParseImageName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Works fine",
			args: args{
				name: "mario_g_marq.png",
			},
			want:    "Mario G Marq",
			wantErr: false,
		}, {
			name: "Should return error",
			args: args{
				name: "mario_g_marq.png.jpq",
			},
			want:    "",
			wantErr: true,
		}, {
			name: "Should return error",
			args: args{
				name: "mario",
			},
			want:    "",
			wantErr: true,
		}, {
			name: "Should return error",
			args: args{
				name: "mario.cpp",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseImageName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseImageName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseImageName() = %v, want %v", got, tt.want)
			}
		})
	}
}
