package proxmox

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCustomStorageDevice_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		want    *CustomStorageDevice
		wantErr bool
	}{
		{
			name: "simple volume",
			line: `"local-lvm:vm-2041-disk-0,discard=on,iothread=1,size=8G"`,
			want: &CustomStorageDevice{
				Discard:    strPtr("on"),
				Enabled:    true,
				FileVolume: "local-lvm:vm-2041-disk-0",
				IOThread:   boolPtr(true),
				Size:       strPtr("8G"),
			},
		},
		{
			name: "raw volume type",
			line: `"nfs:2041/vm-2041-disk-0.raw,discard=ignore,iothread=1,size=8G"`,
			want: &CustomStorageDevice{
				Discard:    strPtr("ignore"),
				Enabled:    true,
				FileVolume: "nfs:2041/vm-2041-disk-0.raw",
				Format:     strPtr("raw"),
				IOThread:   boolPtr(true),
				Size:       strPtr("8G"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CustomStorageDevice{}
			if err := r.UnmarshalJSON([]byte(tt.line)); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			require.Equal(t, tt.want, r)
		})
	}
}

func strPtr(s string) *string {
	return &s
}
func boolPtr(s bool) *CustomBool {
	customBool := CustomBool(s)
	return &customBool
}
