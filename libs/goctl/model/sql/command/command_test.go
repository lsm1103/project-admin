package command

import (
	_ "embed"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/gen"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

var (
	//go:embed testdata/user.sql
	sql string
	cfg = &config.Config{
		NamingFormat: "gozero",
	}
)

func TestFromDDl1(t *testing.T) {
	//err := gen.Clean()
	//assert.Nil(t, err)

	t.Logf("mergeColumns(): %s", mergeColumns(VarStringSliceIgnoreColumns))
	pathx.RegisterGoctlHome("/Users/xm/Desktop/go_package/project-admin/libs/template")
	err := fromDDL(ddlArg{
		src:      "/Users/xm/Desktop/go_package/project-admin/deploy/init.sql",
		dir:      "/Users/xm/Desktop/go_package/project-admin/dataModel/project4",
		cfg:      cfg,
		cache:    true,
		database: "go-zero",
		strict:   false,
		ignoreColumns: mergeColumns([]string{"create_at", "created_at", "create_time", "update_at", "updated_at", "update_time"}),
	})
	t.Logf("err:%+v", err)
}

func TestFromDDL2(t *testing.T) {
	pathx.RegisterGoctlHome("/Users/xm/Desktop/go_package/project-admin/libs/template")
	arg := ddlArg{
		src:      "/Users/xm/Desktop/go_package/project-admin/deploy/init.sql",
		dir:      "/Users/xm/Desktop/go_package/project-admin/dataModel/project11",
		cfg:      cfg,
		cache:    true,
		database: "go-zero",
		strict:   false,
		ignoreColumns: mergeColumns([]string{"create_at", "created_at", "create_time", "update_at", "updated_at", "update_time"}),
	}
	log := console.NewConsole(arg.idea)
	src := strings.TrimSpace(arg.src)
	if len(src) == 0 {
		t.Log("expected path or path globbing patterns, but nothing found")
	}
	files, err := util.MatchFiles(src)
	if err != nil {
		t.Log(err)
	}
	if len(files) == 0 {
		t.Log(errNotMatched)
	}

	generator, err := gen.NewDefaultGenerator(arg.dir, arg.cfg,
		gen.WithConsoleOption(log), gen.WithIgnoreColumns(arg.ignoreColumns))
	if err != nil {
		t.Log(err)
	}

	for _, file := range files {
		err = generator.StartFromDDL(file, arg.cache, arg.strict, arg.database)
		if err != nil {
			t.Log(err)
		}
	}
}


func TestFromDDl(t *testing.T) {
	err := gen.Clean()
	assert.Nil(t, err)

	err = fromDDL(ddlArg{
		src:      "./user.sql",
		dir:      pathx.MustTempDir(),
		cfg:      cfg,
		cache:    true,
		database: "go-zero",
		strict:   false,
	})
	assert.Equal(t, errNotMatched, err)

	// case dir is not exists
	unknownDir := filepath.Join(pathx.MustTempDir(), "test", "user.sql")
	err = fromDDL(ddlArg{
		src:      unknownDir,
		dir:      pathx.MustTempDir(),
		cfg:      cfg,
		cache:    true,
		database: "go_zero",
	})
	assert.True(t, func() bool {
		switch err.(type) {
		case *os.PathError:
			return true
		default:
			return false
		}
	}())

	// case empty src
	err = fromDDL(ddlArg{
		dir:      pathx.MustTempDir(),
		cfg:      cfg,
		cache:    true,
		database: "go_zero",
	})
	if err != nil {
		assert.Equal(t, "expected path or path globbing patterns, but nothing found", err.Error())
	}

	tempDir := filepath.Join(pathx.MustTempDir(), "test")
	err = pathx.MkdirIfNotExist(tempDir)
	if err != nil {
		return
	}

	user1Sql := filepath.Join(tempDir, "user1.sql")
	user2Sql := filepath.Join(tempDir, "user2.sql")

	err = ioutil.WriteFile(user1Sql, []byte(sql), os.ModePerm)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(user2Sql, []byte(sql), os.ModePerm)
	if err != nil {
		return
	}

	_, err = os.Stat(user1Sql)
	assert.Nil(t, err)

	_, err = os.Stat(user2Sql)
	assert.Nil(t, err)

	filename := filepath.Join(tempDir, "usermodel.go")
	fromDDL := func(db string) {
		err = fromDDL(ddlArg{
			src:      filepath.Join(tempDir, "user*.sql"),
			dir:      tempDir,
			cfg:      cfg,
			cache:    true,
			database: db,
		})
		assert.Nil(t, err)

		_, err = os.Stat(filename)
		assert.Nil(t, err)
	}

	fromDDL("go_zero")
	_ = os.Remove(filename)
	fromDDL("go-zero")
	_ = os.Remove(filename)
	fromDDL("1gozero")
}

func Test_parseTableList(t *testing.T) {
	testData := []string{"foo", "b*", "bar", "back_up", "foo,bar,b*"}
	patterns := parseTableList(testData)
	actual := patterns.list()
	expected := []string{"foo", "b*", "bar", "back_up"}
	sort.Slice(actual, func(i, j int) bool {
		return actual[i] > actual[j]
	})
	sort.Slice(expected, func(i, j int) bool {
		return expected[i] > expected[j]
	})
	assert.Equal(t, strings.Join(expected, ","), strings.Join(actual, ","))

	matchTestData := map[string]bool{
		"foo":     true,
		"bar":     true,
		"back_up": true,
		"bit":     true,
		"ab":      false,
		"b":       true,
	}
	for v, expected := range matchTestData {
		actual := patterns.Match(v)
		assert.Equal(t, expected, actual)
	}
}
