package cmd

import "testing"

func TestIsSafePathExcluded(t *testing.T) {
	exactPaths := map[string]struct{}{
		"/unsubscribe_new.html": {},
		"/api/unsubscribe_new":  {},
	}

	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "exact unsubscribe page",
			path: "/unsubscribe_new.html",
			want: true,
		},
		{
			name: "exact unsubscribe api",
			path: "/api/unsubscribe_new",
			want: true,
		},
		{
			name: "pmta tracked link",
			path: "/pmta/encrypted-token",
			want: true,
		},
		{
			name: "pmta root without trailing slash",
			path: "/pmta",
			want: false,
		},
		{
			name: "roundcube root without trailing slash",
			path: "/roundcube",
			want: true,
		},
		{
			name: "roundcube root with trailing slash",
			path: "/roundcube/",
			want: true,
		},
		{
			name: "roundcube child path",
			path: "/roundcube/index.php",
			want: true,
		},
		{
			name: "protected backend path",
			path: "/settings",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSafePathExcluded(tt.path, exactPaths)
			if got != tt.want {
				t.Fatalf("isSafePathExcluded(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}
