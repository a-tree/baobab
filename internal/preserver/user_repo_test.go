package preserver

import (
	"baobab/internal/user"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func Test_UserMapping_Strict(t *testing.T) {
	src := &user.User{
		Name:  "Gopher",
		Email: "gopher@example.com",
	}
	dst, _ := UserRepoMapToDB(src)

	// 期待する DB モデルの状態
	expected := &UserDB{
		Name:  "Gopher",
		Email: "gopher@example.com",
	}

	// 差分があればテストを落とす
	// もし UserDB 側の修正を忘れると、ここで差分が出て検知できる
	diff := cmp.Diff(expected, dst, cmpopts.IgnoreFields(UserDB{}, "Model"))
	assert.Equal(t, diff, "")
	if diff != "" {
		t.Errorf("Mapping mismatch (-want +got):\n%s", diff)
	}
}
