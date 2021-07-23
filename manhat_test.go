package manhat

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCalculateMatrixSize(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name   string
		number int
		want   int
	}{
		{
			name:   "One",
			number: 1,
			want:   1,
		},
		{
			name:   "Inside 3x3",
			number: 8,
			want:   3,
		},
		{
			name:   "Highest in 3x3",
			number: 9,
			want:   3,
		},
		{
			name:   "Over the edge of 3x3",
			number: 10,
			want:   5,
		},
		{
			name:   "End of 5x5",
			number: 25,
			want:   5,
		},
		{
			name:   "One beyond 5x5",
			number: 26,
			want:   7,
		},
		{
			name:   "Max in 9x9",
			number: 81,
			want:   9,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculateMatrixSize(tc.number)
			if err != nil {
				t.Fatalf("calculateMatrixSize(%v) got err %v", tc.number, err)
			}
			if got != tc.want {
				t.Errorf("calculateMatrixSize(%v) got %v, want %v\n", tc.number, got, tc.want)
			}
		})
	}
}

func TestCalculateMatricSizeInvalid(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name   string
		number int
	}{
		{
			name:   "Zero number",
			number: 0,
		},
		{
			name:   "Less than zero",
			number: -1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := calculateMatrixSize(tc.number)
			if err == nil {
				t.Fatalf("calculateMatrixSize(%v) expected err, did not get one", tc.number)
			}
		})
	}
}

func TestBuildSpiralMatrix(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		size int
		want [][]int
	}{
		{
			name: "Min size",
			size: 1,
			want: [][]int{
				{1},
			},
		},
		{
			name: "Size 3x3",
			size: 3,
			want: [][]int{
				{5, 4, 3},
				{6, 1, 2},
				{7, 8, 9},
			},
		},
		{
			name: "Size 5x5",
			size: 5,
			want: [][]int{
				{17, 16, 15, 14, 13},
				{18, 5, 4, 3, 12},
				{19, 6, 1, 2, 11},
				{20, 7, 8, 9, 10},
				{21, 22, 23, 24, 25},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, _, err := buildSpiralMatrix(tc.size)
			if err != nil {
				t.Fatalf("buildSpiralMatrix(%v) got err %v", tc.size, err)
			}

			if !cmp.Equal(tc.want, got) {
				t.Errorf("buildSpiratMatrix(%v) got \n%v\n", tc.size, cmp.Diff(tc.want, got))
			}
		})
	}
}

func TestBuildSpiralMatrixInvalid(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		size int
	}{
		{
			name: "Zero size",
			size: 0,
		},
		{
			name: "Less than zero size",
			size: -1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, _, err := buildSpiralMatrix(tc.size)
			if err == nil {
				t.Fatalf("buildSpiralMatrix(%v) did not get error", tc.size)
			}
		})
	}

}

func TestGenerateNumbers(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name   string
		size   int
		number int
		want   []int
	}{
		{
			name:   "Min size",
			size:   1,
			number: 1,
			want:   []int{0, 0},
		},
		{
			name:   "Size 3x3",
			size:   3,
			number: 9,
			want:   []int{2, 2},
		},
		{
			name:   "Size 3x3",
			size:   3,
			number: 7,
			want:   []int{0, 2},
		},
		{
			name:   "Size 3x3",
			size:   3,
			number: 2,
			want:   []int{2, 1},
		},
		{
			name:   "Size 5x5",
			size:   5,
			number: 14,
			want:   []int{3, 0},
		},
		{
			name:   "Size 5x5",
			size:   5,
			number: 20,
			want:   []int{0, 3},
		},
		{
			name:   "Size 5x5",
			size:   5,
			number: 4,
			want:   []int{2, 1},
		},
		{
			name:   "Size 5x5",
			size:   5,
			number: 6,
			want:   []int{1, 2},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, coordinates, err := buildSpiralMatrix(tc.size)

			if err != nil {
				t.Fatalf("buildSpiralMatrix(%v) for number %v got err %v", tc.size, tc.number, err)
			}

			got := coordinates[tc.number]

			if !cmp.Equal(tc.want, got) {
				t.Errorf("buildSpiratMatrix(%v) for number %v, got error \n%v\n", tc.size, tc.number, cmp.Diff(tc.want, got))
			}
		})
	}
}

func TestCalculateDistance(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name         string
		location     int
		wantDistance int
	}{
		{
			name:         "Location 1",
			location:     1,
			wantDistance: 0,
		},
		{
			name:         "Location 12",
			location:     12,
			wantDistance: 3,
		},
		{
			name:         "Location 12",
			location:     23,
			wantDistance: 2,
		},
		{
			name:         "Location 1024",
			location:     1024,
			wantDistance: 31,
		},
		{
			name:         "Location 368078",
			location:     368078,
			wantDistance: 371,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			got, err := CalculateDistance(tc.location)
			if err != nil {
				t.Fatalf("CalculateDistance(%v) got error %v", tc.location, err)
			}

			if tc.wantDistance != got {
				t.Errorf("CalculateDistance(%v) got %v, want %v", tc.location, int(got), tc.wantDistance)
			}
		})
	}
}

func TestCalculateDistanceInvalid(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name         string
		location     int
		wantDistance int
	}{
		{
			name:         "Invalid input",
			location:     0,
			wantDistance: 0,
		},
		{
			name:         "Ivalid input",
			location:     -1,
			wantDistance: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			got, err := CalculateDistance(tc.location)
			if err == nil {
				t.Fatalf("CalculateDistance(%v) got error %v", tc.location, err)
			}

			if tc.wantDistance != got {
				t.Errorf("CalculateDistance(%v) got %v, want %v", tc.location, got, tc.wantDistance)
			}
		})
	}
}
