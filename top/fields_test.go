package metricsTop

import (
	"testing"
)

func TestFields(t *testing.T) {
	s := `29525 ubuntu    20   0  287848  50196  11632 S  6.7  1.2   1:12.45 /usr/bin/python3 manage.py runserver 0.0.0.0:8888 --insecure`
	expected := []string{
		"29525",
		"ubuntu",
		"20",
		"0",
		"287848",
		"50196",
		"11632",
		"S",
		"6.7",
		"1.2",
		"1:12.45",
		"/usr/bin/python3 manage.py runserver 0.0.0.0:8888 --insecure",
	}
	got := Fields(s)

	if len(expected) != len(got) {
		t.Errorf("len not equal, expected %d got %d", len(expected), len(got))
	}

	for idx, value := range got {
		if expected[idx] != value {
			t.Errorf("idx %d expected %s got %s", idx, expected[idx], value)
		}
	}

}
