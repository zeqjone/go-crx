package crx

import (
	"testing"
)

// func TestUnpack(t *testing.T) {
// 	stats := map[string]int{
// 		"background.js":     0,
// 		"images":            0,
// 		"images/image.jpeg": 0,
// 		"manifest.json":     0,
// 	}
// 	err := Unpack("./testdata/unpack/extension.crx")
// 	assert.Nil(t, err)
// 	err = filepath.Walk("./testdata/unpack/extension", func(path string, info os.FileInfo, err error) error {
// 		relpath, err := filepath.Rel("./testdata/unpack/extension", path)
// 		assert.Nil(t, err)
// 		stats[relpath]++
// 		return nil
// 	})
// 	assert.Nil(t, err)
// 	for _, v := range stats {
// 		assert.Equal(t, 1, v)
// 	}
// 	err = os.RemoveAll("./testdata/unpack/extension")
// 	assert.Nil(t, err)
// }

func TestUnpack(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "crx2",
			args: args{
				filename: "./testdata/crx2/thy_4.1.crx",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unpack(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("Unpack() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
