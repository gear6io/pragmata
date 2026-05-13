package pqlgrammar_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gear6io/pragmata/pkg/grammars/pqlgrammar"
)

// ── Test helpers ───────────────────────────────────────────────────────────────

type pqlErrListener struct {
	*antlr.DefaultErrorListener
	errs []string
}

func (l *pqlErrListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, col int, msg string, _ antlr.RecognitionException) {
	l.errs = append(l.errs, fmt.Sprintf("line %d:%d %s", line, col, msg))
}

func parsePQL(input string) (pqlgrammar.IPipelineContext, []string) {
	el := &pqlErrListener{}

	lexer := pqlgrammar.NewPQLLexer(antlr.NewInputStream(input))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := pqlgrammar.NewPQL(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(el)

	return p.Pipeline(), el.errs
}

func mustParsePQL(t *testing.T, input string) pqlgrammar.IPipelineContext {
	t.Helper()
	tree, errs := parsePQL(input)
	if len(errs) > 0 {
		t.Fatalf("unexpected parse errors:\n%s", strings.Join(errs, "\n"))
	}
	return tree
}

// fromOnly constructs a minimal valid pipeline — useful when a test only cares
// about a single non-from transform and needs a preceding from.
func fromOnly(transform string) string {
	return "from orders\n" + transform
}

// ── Valid inputs ───────────────────────────────────────────────────────────────

func TestPQL_Valid(t *testing.T) {
	cases := []struct {
		name  string
		input string
	}{
		// from
		{"from simple", "from orders\n"},
		{"from final", "from orders final\n"},
		{"from mixed case", "from orders FINAL\n"},

		// filter — content captured as opaque FILTER_LINE tokens
		{"filter single line", "from orders\nfilter status = 200\n"},
		{"filter bare word", "from orders\nfilter env = prod\n"},
		{"filter with and", "from orders\nfilter status = 200 AND env = prod\n"},
		{"filter multiline continuation", "from orders\nfilter status = 200\n  AND env = prod\n"},

		// derive
		{"derive single", fromOnly("derive { day = toStartOfDay(ts) }\n")},
		{"derive two cols", fromOnly("derive { day = toStartOfDay(ts), tier = 'low' }\n")},
		{"derive nested parens", fromOnly("derive { tier = multiIf(amount > 1000, 'high', 'low') }\n")},
		{"derive double call", fromOnly("derive { p95 = quantile(0.95)(latency) }\n")},
		{"derive arithmetic", fromOnly("derive { tax = amount * 0.1 }\n")},
		{"derive cast expr", fromOnly("derive { n = attributes.status::Int64 }\n")},

		// select
		{"select cols", fromOnly("select { user_id, amount }\n")},
		{"select alias", fromOnly("select { total = sum(amount) }\n")},
		{"select mixed", fromOnly("select { user_id, total = sum(amount) }\n")},

		// group + aggregate
		{"group single key", fromOnly("group { country } ( aggregate { total = sum(amount) } )\n")},
		{"group multi key", fromOnly("group { day, country } ( aggregate { total = sum(amount), n = count() } )\n")},
		{"group computed key", fromOnly("group { day = toStartOfDay(ts) } ( aggregate { total = sum(amount) } )\n")},
		{"aggregate uniq", fromOnly("group { country } ( aggregate { users = uniqExact(user_id) } )\n")},
		{"aggregate quantile", fromOnly("group { country } ( aggregate { p95 = quantile(0.95)(latency) } )\n")},

		// join
		{"join left", fromOnly("join side:left dim_country (country = dim_country.code)\n")},
		{"join right", fromOnly("join side:right sessions (user_id = sessions.uid)\n")},
		{"join inner", fromOnly("join side:inner lookup (id = lookup.key)\n")},
		{"join full", fromOnly("join side:full other (id = other.id)\n")},
		{"join self shorthand", fromOnly("join side:left dim (==user_id)\n")},
		{"join no side", fromOnly("join other (id = other.id)\n")},

		// array_join
		{"array_join plain", fromOnly("array_join tags\n")},
		{"array_join alias", fromOnly("array_join tags as t\n")},

		// sort
		{"sort desc", fromOnly("sort { -amount }\n")},
		{"sort asc implicit", fromOnly("sort { amount }\n")},
		{"sort asc explicit", fromOnly("sort { +amount }\n")},
		{"sort multi", fromOnly("sort { -day, revenue }\n")},
		{"sort three", fromOnly("sort { -day, -revenue, +user_id }\n")},

		// take / skip
		{"take simple", fromOnly("take 10\n")},
		{"take range", fromOnly("take 10..20\n")},
		{"take large", fromOnly("take 1000\n")},
		{"skip", fromOnly("skip 5\n")},
		{"take then skip", fromOnly("take 100\nskip 20\n")},

		// window
		{"window rank", fromOnly("window { rn = rank() OVER (PARTITION BY country ORDER BY ts) }\n")},
		{"window running total", fromOnly("window { rt = sum(amount) OVER (PARTITION BY country ORDER BY ts) }\n")},

		// multi-transform pipelines
		{"from filter", "from orders\nfilter env = prod\n"},
		{"from derive", "from orders\nderive { day = toStartOfDay(ts) }\n"},
		{"from filter derive group sort take",
			"from orders\n" +
				"filter status = 200\n" +
				"derive { day = toStartOfDay(ts) }\n" +
				"group { day, country } ( aggregate { revenue = sum(amount) } )\n" +
				"sort { -day }\n" +
				"take 100\n"},
		{"from filter select", "from orders\nfilter env = prod\nselect { user_id, amount }\n"},
		{"from group sort", "from orders\ngroup { country } ( aggregate { total = sum(amount) } )\nsort { -total }\n"},
		{"from join filter",
			"from orders\njoin side:left regions (country = regions.code)\nfilter revenue > 1000\n"},
		{"from array_join group",
			"from orders\narray_join tags\ngroup { tags } ( aggregate { n = count() } )\n"},
		{"from final filter",
			"from orders final\nfilter env = prod AND status = 200\n"},

		// filter multiline with next transform
		{"filter multi then derive",
			"from orders\n" +
				"filter status = 200\n" +
				"  AND env = prod\n" +
				"derive { day = toStartOfDay(ts) }\n"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, errs := parsePQL(tc.input)
			if len(errs) > 0 {
				t.Errorf("unexpected errors:\n%s", strings.Join(errs, "\n"))
			}
		})
	}
}

// ── Invalid inputs ─────────────────────────────────────────────────────────────

func TestPQL_Invalid(t *testing.T) {
	cases := []struct {
		name  string
		input string
	}{
		{"empty input", ""},
		{"from without table", "from\n"},
		{"derive without braces", "from orders\nderive day = toStartOfDay(ts)\n"},
		{"derive empty braces", "from orders\nderive {}\n"},
		{"group without aggregate", "from orders\ngroup { country }\n"},
		{"take without number", "from orders\ntake\n"},
		{"skip without number", "from orders\nskip\n"},
		{"sort empty braces", "from orders\nsort {}\n"},
		{"select empty braces", "from orders\nselect {}\n"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, errs := parsePQL(tc.input)
			if len(errs) == 0 {
				t.Errorf("expected parse errors for %q but got none", tc.input)
			}
		})
	}
}

// ── Structural assertions ──────────────────────────────────────────────────────

func TestPQL_Pipeline_TransformCount(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{"from only", "from orders\n", 1},
		{"from filter", "from orders\nfilter env = prod\n", 2},
		{"from derive group sort take",
			"from orders\n" +
				"derive { day = toStartOfDay(ts) }\n" +
				"group { day } ( aggregate { n = count() } )\n" +
				"sort { -day }\n" +
				"take 10\n",
			5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tree := mustParsePQL(t, tc.input)
			if got := len(tree.AllTransform()); got != tc.want {
				t.Errorf("transform count = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestPQL_From_TableName(t *testing.T) {
	cases := []struct{ table string }{
		{"orders"},
		{"events"},
		{"dim_country"},
		{"my_datasource"},
	}

	for _, tc := range cases {
		t.Run(tc.table, func(t *testing.T) {
			tree := mustParsePQL(t, "from "+tc.table+"\n")
			from := tree.Transform(0).FromTransform()
			if from == nil {
				t.Fatal("expected fromTransform")
			}
			if got := from.IDENT().GetText(); got != tc.table {
				t.Errorf("table name = %q, want %q", got, tc.table)
			}
		})
	}
}

func TestPQL_From_Final(t *testing.T) {
	t.Run("without final", func(t *testing.T) {
		tree := mustParsePQL(t, "from orders\n")
		from := tree.Transform(0).FromTransform()
		if from.KW_FINAL() != nil {
			t.Error("expected KW_FINAL to be absent")
		}
	})

	t.Run("with final", func(t *testing.T) {
		tree := mustParsePQL(t, "from orders final\n")
		from := tree.Transform(0).FromTransform()
		if from.KW_FINAL() == nil {
			t.Error("expected KW_FINAL to be present")
		}
	})
}

func TestPQL_Filter_Body(t *testing.T) {
	input := "from orders\nfilter status = 200\n"
	tree := mustParsePQL(t, input)

	filterT := tree.Transform(1).FilterTransform()
	if filterT == nil {
		t.Fatal("expected filterTransform at index 1")
	}

	lines := filterT.FilterBody().AllFILTER_LINE()
	if len(lines) == 0 {
		t.Fatal("expected at least one FILTER_LINE token")
	}
	// the captured line includes the trailing newline
	got := strings.TrimRight(lines[0].GetText(), "\r\n")
	if got != "status = 200" {
		t.Errorf("filter line = %q, want %q", got, "status = 200")
	}
}

func TestPQL_Filter_MultilineContinuation(t *testing.T) {
	input := "from orders\nfilter status = 200\n  AND env = prod\nderive { day = toStartOfDay(ts) }\n"
	tree := mustParsePQL(t, input)

	filterT := tree.Transform(1).FilterTransform()
	if filterT == nil {
		t.Fatal("expected filterTransform at index 1")
	}
	// two FILTER_LINE tokens: one for each input line before derive
	if got := len(filterT.FilterBody().AllFILTER_LINE()); got != 2 {
		t.Errorf("FILTER_LINE count = %d, want 2", got)
	}
	// third transform must be derive (filter body did not consume it)
	if tree.Transform(2).DeriveTransform() == nil {
		t.Error("expected deriveTransform at index 2")
	}
}

func TestPQL_Derive_AssignmentNames(t *testing.T) {
	input := "from orders\nderive { day = toStartOfDay(ts), tier = 'low' }\n"
	tree := mustParsePQL(t, input)

	derive := tree.Transform(1).DeriveTransform()
	if derive == nil {
		t.Fatal("expected deriveTransform")
	}

	assignments := derive.AssignmentList().AllAssignment()
	if len(assignments) != 2 {
		t.Fatalf("assignment count = %d, want 2", len(assignments))
	}

	names := []string{
		assignments[0].IDENT().GetText(),
		assignments[1].IDENT().GetText(),
	}
	if names[0] != "day" {
		t.Errorf("assignment[0] name = %q, want %q", names[0], "day")
	}
	if names[1] != "tier" {
		t.Errorf("assignment[1] name = %q, want %q", names[1], "tier")
	}
}

func TestPQL_Derive_NestedParens(t *testing.T) {
	// multiIf has commas inside parens — must not be split into multiple assignments
	input := "from orders\nderive { tier = multiIf(amount > 1000, 'high', 'low') }\n"
	tree := mustParsePQL(t, input)

	derive := tree.Transform(1).DeriveTransform()
	assignments := derive.AssignmentList().AllAssignment()
	if len(assignments) != 1 {
		t.Errorf("assignment count = %d, want 1 (nested commas must not split)", len(assignments))
	}
}

func TestPQL_Group_Keys(t *testing.T) {
	input := "from orders\ngroup { day, country } ( aggregate { total = sum(amount) } )\n"
	tree := mustParsePQL(t, input)

	group := tree.Transform(1).GroupTransform()
	if group == nil {
		t.Fatal("expected groupTransform")
	}

	keys := group.KeyList().AllKeyItem()
	if len(keys) != 2 {
		t.Fatalf("key count = %d, want 2", len(keys))
	}
}

func TestPQL_Sort_Direction(t *testing.T) {
	input := "from orders\nsort { -day, revenue, +code }\n"
	tree := mustParsePQL(t, input)

	sort := tree.Transform(1).SortTransform()
	if sort == nil {
		t.Fatal("expected sortTransform")
	}

	items := sort.SortList().AllSortItem()
	if len(items) != 3 {
		t.Fatalf("sort item count = %d, want 3", len(items))
	}

	if _, ok := items[0].(*pqlgrammar.DescSortContext); !ok {
		t.Errorf("items[0]: expected DescSortContext (-day), got %T", items[0])
	}
	if _, ok := items[1].(*pqlgrammar.AscSortContext); !ok {
		t.Errorf("items[1]: expected AscSortContext (revenue), got %T", items[1])
	}
	if _, ok := items[2].(*pqlgrammar.AscSortExplicitContext); !ok {
		t.Errorf("items[2]: expected AscSortExplicitContext (+code), got %T", items[2])
	}
}

func TestPQL_Take_Simple(t *testing.T) {
	input := "from orders\ntake 50\n"
	tree := mustParsePQL(t, input)

	take := tree.Transform(1).TakeTransform()
	if take == nil {
		t.Fatal("expected takeTransform")
	}
	if take.RANGE() != nil {
		t.Error("expected no RANGE token for simple take")
	}
	if got := take.INTEGER(0).GetText(); got != "50" {
		t.Errorf("take n = %q, want %q", got, "50")
	}
}

func TestPQL_Take_Range(t *testing.T) {
	input := "from orders\ntake 10..20\n"
	tree := mustParsePQL(t, input)

	take := tree.Transform(1).TakeTransform()
	if take == nil {
		t.Fatal("expected takeTransform")
	}
	if take.RANGE() == nil {
		t.Error("expected RANGE token for range take")
	}
	if got := take.INTEGER(0).GetText(); got != "10" {
		t.Errorf("take start = %q, want %q", got, "10")
	}
	if got := take.INTEGER(1).GetText(); got != "20" {
		t.Errorf("take end = %q, want %q", got, "20")
	}
}

func TestPQL_Join_Side(t *testing.T) {
	cases := []struct {
		side  string
		check func(pqlgrammar.IJoinSideContext) bool
	}{
		{"left", func(s pqlgrammar.IJoinSideContext) bool { return s.KW_LEFT() != nil }},
		{"right", func(s pqlgrammar.IJoinSideContext) bool { return s.KW_RIGHT() != nil }},
		{"inner", func(s pqlgrammar.IJoinSideContext) bool { return s.KW_INNER() != nil }},
		{"full", func(s pqlgrammar.IJoinSideContext) bool { return s.KW_FULL() != nil }},
	}

	for _, tc := range cases {
		t.Run(tc.side, func(t *testing.T) {
			input := fmt.Sprintf("from orders\njoin side:%s dim (id = dim.key)\n", tc.side)
			tree := mustParsePQL(t, input)
			join := tree.Transform(1).JoinTransform()
			if join == nil {
				t.Fatal("expected joinTransform")
			}
			if !tc.check(join.JoinSide()) {
				t.Errorf("expected KW_%s in join side", strings.ToUpper(tc.side))
			}
		})
	}
}

func TestPQL_Join_SelfShorthand(t *testing.T) {
	input := "from orders\njoin side:left dim (==user_id)\n"
	tree := mustParsePQL(t, input)

	join := tree.Transform(1).JoinTransform()
	if join == nil {
		t.Fatal("expected joinTransform")
	}
	if _, ok := join.JoinCond().(*pqlgrammar.SelfJoinCondContext); !ok {
		t.Errorf("expected SelfJoinCondContext, got %T", join.JoinCond())
	}
}

func TestPQL_ArrayJoin_WithAlias(t *testing.T) {
	input := "from orders\narray_join tags as t\n"
	tree := mustParsePQL(t, input)

	aj := tree.Transform(1).ArrayJoinTransform()
	if aj == nil {
		t.Fatal("expected arrayJoinTransform")
	}
	idents := aj.AllIDENT()
	if len(idents) != 2 {
		t.Fatalf("array_join ident count = %d, want 2 (column + alias)", len(idents))
	}
	if got := idents[0].GetText(); got != "tags" {
		t.Errorf("column = %q, want %q", got, "tags")
	}
	if got := idents[1].GetText(); got != "t" {
		t.Errorf("alias = %q, want %q", got, "t")
	}
}

func TestPQL_FullPipeline(t *testing.T) {
	input := "" +
		"from orders\n" +
		"filter status = 200\n" +
		"  AND env = prod\n" +
		"derive { day = toStartOfDay(observed_timestamp) }\n" +
		"group { day, country } ( aggregate { revenue = sum(amount), orders = count() } )\n" +
		"sort { -day, revenue }\n" +
		"take 100\n"

	tree := mustParsePQL(t, input)

	transforms := tree.AllTransform()
	if len(transforms) != 6 {
		t.Fatalf("transform count = %d, want 6", len(transforms))
	}

	if transforms[0].FromTransform() == nil {
		t.Error("transform[0]: expected fromTransform")
	}
	if transforms[1].FilterTransform() == nil {
		t.Error("transform[1]: expected filterTransform")
	}
	if transforms[2].DeriveTransform() == nil {
		t.Error("transform[2]: expected deriveTransform")
	}
	if transforms[3].GroupTransform() == nil {
		t.Error("transform[3]: expected groupTransform")
	}
	if transforms[4].SortTransform() == nil {
		t.Error("transform[4]: expected sortTransform")
	}
	if transforms[5].TakeTransform() == nil {
		t.Error("transform[5]: expected takeTransform")
	}
}
