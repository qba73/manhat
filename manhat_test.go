package manhat

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCalculateMatrixSize(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		number      float64
		want        float64
		expectedErr bool
	}{
		{
			name:        "Zero number",
			number:      0,
			want:        0,
			expectedErr: true,
		},
		{
			name:        "Less than zero",
			number:      -1,
			want:        0,
			expectedErr: true,
		},
		{
			name:        "One",
			number:      1,
			want:        1,
			expectedErr: false,
		},
		{
			name:        "Inside 3x3",
			number:      8,
			want:        3,
			expectedErr: false,
		},
		{
			name:        "Highest in 3x3",
			number:      9,
			want:        3,
			expectedErr: false,
		},
		{
			name:        "Over the edge of 3x3",
			number:      10,
			want:        5,
			expectedErr: false,
		},
		{
			name:        "End of 5x5",
			number:      25,
			want:        5,
			expectedErr: false,
		},
		{
			name:        "One beyond 5x5",
			number:      26,
			want:        7,
			expectedErr: false,
		},
		{
			name:        "Max in 9x9",
			number:      81,
			want:        9,
			expectedErr: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculateMatrixSize(tc.number)
			if (err != nil) && !tc.expectedErr {
				t.Fatalf("%s, calculateMatrixSize(%v) got err %v", tc.name, tc.number, err)
			}
			if !tc.expectedErr && (got != tc.want) {
				t.Errorf("%s, calculateMatrixSize(%v) got %v, want %v\n", tc.name, tc.number, got, tc.want)
			}
		})
	}
}

func TestBuildSpiralMatrix(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		size        int
		want        [][]int
		expectedErr bool
	}{
		{
			name: "Min size",
			size: 1,
			want: [][]int{
				{1},
			},
			expectedErr: false,
		},
		{
			name: "Size 3x3",
			size: 3,
			want: [][]int{
				{5, 4, 3},
				{6, 1, 2},
				{7, 8, 9},
			},
			expectedErr: false,
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
			expectedErr: false,
		},
		{
			name:        "Zero size",
			size:        0,
			want:        [][]int{{}},
			expectedErr: true,
		},
		{
			name:        "Less than zero size",
			size:        -1,
			want:        [][]int{{}},
			expectedErr: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, _, err := buildSpiralMatrix(tc.size)
			if (err != nil) && !tc.expectedErr {
				t.Fatalf("%s, buildSpiralMatrix(%v) got err %v", tc.name, tc.size, err)
			}

			if !tc.expectedErr && !cmp.Equal(tc.want, got) {
				t.Errorf("%s, buildSpiratMatrix(%v) got error \n%v\n", tc.name, tc.size, cmp.Diff(tc.want, got))
			}
		})
	}
}

func TestGenerateNumbers(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		size        int
		number      int
		want        []int
		expectedErr bool
	}{
		{
			name:        "Min size",
			size:        1,
			number:      1,
			want:        []int{0, 0},
			expectedErr: false,
		},
		{
			name:        "Size 3x3",
			size:        3,
			number:      9,
			want:        []int{2, 2},
			expectedErr: false,
		},
		{
			name:        "Size 3x3",
			size:        3,
			number:      7,
			want:        []int{0, 2},
			expectedErr: false,
		},
		{
			name:        "Size 3x3",
			size:        3,
			number:      2,
			want:        []int{2, 1},
			expectedErr: false,
		},
		{
			name:        "Size 5x5",
			size:        5,
			number:      14,
			want:        []int{3, 0},
			expectedErr: false,
		},
		{
			name:        "Size 5x5",
			size:        5,
			number:      20,
			want:        []int{0, 3},
			expectedErr: false,
		},
		{
			name:        "Size 5x5",
			size:        5,
			number:      4,
			want:        []int{2, 1},
			expectedErr: false,
		},
		{
			name:        "Size 5x5",
			size:        5,
			number:      6,
			want:        []int{1, 2},
			expectedErr: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, coordinates, err := buildSpiralMatrix(tc.size)

			if (err != nil) && !tc.expectedErr {
				t.Fatalf("%s, buildSpiralMatrix(%v) for number %v got err %v", tc.name, tc.size, tc.number, err)
			}

			got := coordinates[tc.number]

			if !tc.expectedErr && !cmp.Equal(tc.want, got) {
				t.Errorf("%s, buildSpiratMatrix(%v) for number %v, got error \n%v\n", tc.name, tc.size, tc.number, cmp.Diff(tc.want, got))
			}
		})
	}
}

func TestManhattanDistanceFormula(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name   string
		center point
		poi    point
		want   float64
	}{
		{
			name:   "",
			center: point{x: 1, y: 1},
			poi:    point{x: 1, y: 1},
			want:   0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}

func TestCalculateDistance(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name         string
		location     int
		wantDistance int
		expectedErr  bool
	}{
		{
			name:         "Location 1",
			location:     1,
			wantDistance: 0,
			expectedErr:  false,
		},
		{
			name:         "Location 12",
			location:     12,
			wantDistance: 3,
			expectedErr:  false,
		},
		{
			name:         "Location 12",
			location:     23,
			wantDistance: 2,
			expectedErr:  false,
		},
		{
			name:         "Location 1024",
			location:     1024,
			wantDistance: 31,
			expectedErr:  false,
		},
		{
			name:         "Location 368078",
			location:     368078,
			wantDistance: 371,
			expectedErr:  false,
		},
		{
			name:         "Invalid input",
			location:     0,
			wantDistance: 0,
			expectedErr:  true,
		},
		{
			name:         "Ivalid input",
			location:     -1,
			wantDistance: 0,
			expectedErr:  true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			got, err := CalculateDistance(float64(tc.location))
			if (err != nil) && !tc.expectedErr {
				t.Fatalf("%s, CalculateDistance(%v) got error %v", tc.name, tc.location, err)
			}

			if !tc.expectedErr && (tc.wantDistance != int(got)) {
				t.Errorf("%s, CalculateDistance(%v) got %v, want %v", tc.name, tc.location, int(got), tc.wantDistance)
			}
		})
	}
}
