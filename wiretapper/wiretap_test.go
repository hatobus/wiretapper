package wiretapper

import "testing"

func Add2(a, b int64) int64 {
	return a+b
}

func CheckCallCountAdd2(wt *WireTapper, a, b int64) int64 {
	wt.Countup()

	return Add2(a, b)
}

func TestWireTap(t *testing.T) {

	type testCase struct {
		wt *WireTapper
		a int64
		b int64
		iteratorCount int
		wantCallCount uint64
		wantAnswer int64
	}

	testCases := map[string]testCase{
		"1回だけ関数が呼ばれる": {
			wt: Initialize(),
			a: 1,
			b: 1,
			iteratorCount: 1,
			wantCallCount: 1,
			wantAnswer: 2,
		},
		"2回関数が呼ばれる": {
			wt: Initialize(),
			a: 1,
			b: 1,
			iteratorCount: 2,
			wantCallCount: 2,
			wantAnswer: 4,
		},
		"3回関数が呼ばれる": {
			wt: Initialize(),
			a: 1,
			b: 1,
			iteratorCount: 3,
			wantCallCount: 3,
			wantAnswer: 6,
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			var ans int64

			for i := 0; i<tc.iteratorCount; i ++ {
				ans += CheckCallCountAdd2(tc.wt, tc.a, tc.b)
			}

			if ans != tc.wantAnswer {
				t.Fatalf("invalid answer: want %v, got %v", tc.wantAnswer, ans)
			}

			callCount := tc.wt.GetCounter()
			if callCount != tc.wantCallCount {
				t.Fatalf("invalid callCount: want %v, got %v", tc.wantCallCount, callCount)
			}
		})
	}

}
