package goignore

import (
	"reflect"
	"sort"
	"testing"
)

func TestTemplates_IsSupportedTemplates(t *testing.T) {
	templates := &Templates{}

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"android", []string{"android"}, false},
		{"phpstorm+iml", []string{"phpstorm+iml"}, false},
		{"objective-c", []string{"objective-c"}, false},
		{"android,phpstorm+iml,objective-c", []string{"android", "phpstorm+iml", "objective-c"}, false},
		{"1c,1c-bitrix,a-frame,zukencr8000", []string{"1c", "1c-bitrix", "a-frame", "zukencr8000"}, false},
		{"*/+0", []string{"*/+0"}, true},
		{"dk", []string{"dk"}, true},
		{"+-*/& ,dk", []string{"+-*/& ", "dk"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := templates.IsSupportedTemplates(tt.args...); (err != nil) != tt.wantErr {
				t.Errorf("Templates.IsSupportedTemplates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTemplates_IsCustomTemplate(t *testing.T) {
	tests := []struct {
		name      string
		templates Templates
		args      string
		wantErr   bool
	}{
		{"Android",
			Templates{
				CustomTemplates: map[string]string{
					"Android": "android,linux,intellij",
				},
			},
			"@Android",
			false,
		},
		{"golang",
			Templates{
				CustomTemplates: map[string]string{
					"Android": "android,linux,intellij",
				},
			},
			"@golang",
			true,
		},
		{"golang",
			Templates{
				CustomTemplates: map[string]string{
					"Android": "android,linux,intellij",
				},
			},
			"golang",
			true,
		},
		{"golang - empty templates",
			Templates{},
			"golang",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.templates.IsCustomTemplate(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Templates.IsCustomTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTemplates_PreprocessInputTemplates(t *testing.T) {

	type args struct {
		inputTemplates []string
	}
	tests := []struct {
		name      string
		templates Templates
		args      []string
		want      []string
	}{
		{"go",
			Templates{},
			[]string{"go"},
			[]string{"go"},
		},
		{"go,linux",
			Templates{},
			[]string{"go", "linux"},
			[]string{"go", "linux"},
		},
		{"@Python",
			Templates{
				[]string{},
				map[string]string{"Python": "python,linux"},
			},
			[]string{"@Python"},
			[]string{"@Python"},
		},
		{"@Python,@Go",
			Templates{
				[]string{},
				map[string]string{"Python": "python,linux", "Go": "go,linux"},
			},
			[]string{"@Python", "@Go"},
			[]string{"python", "go", "linux"},
		},
		{"@Python,jetbrains,@Go",
			Templates{
				[]string{},
				map[string]string{"Python": "python,linux", "Go": "go,linux"},
			},
			[]string{"@Python", "jetbrains", "@Go"},
			[]string{"python", "go", "linux", "jetbrains"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.templates.PreprocessInputTemplates(tt.args...)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Templates.PreprocessInputTemplates() = %v, want %v", got, tt.want)
			}
		})
	}
}
