package preserver

import (
	"baobab/internal/place"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func Test_PlaceMapping_Strict(t *testing.T) {
	src := &place.Place{
		Address:    "Tokyo Station, 1-chome Marunouchi",
		Country:    "Japan",
		Prefecture: "Tokyo",
		City:       "Chiyoda-ku",
		Postal:     "100-0005",
	}
	dst, _ := PlaceRepoMapToDB(src)

	// 期待する DB モデルの状態
	expected := &PlaceDB{
		Address:    "Tokyo Station, 1-chome Marunouchi",
		Country:    "Japan",
		Prefecture: "Tokyo",
		City:       "Chiyoda-ku",
		Postal:     "100-0005",
	}

	// 差分があればテストを落とす
	// もし PlaceDB 側の修正を忘れると、ここで差分が出て検知できる
	diff := cmp.Diff(expected, dst, cmpopts.IgnoreFields(PlaceDB{}, "Model"))
	assert.Equal(t, diff, "")
	if diff != "" {
		t.Errorf("Mapping mismatch (-want +got):\n%s", diff)
	}
}
