package postgres

import "testing"

func sp(s string) *string { return &s }

func TestBuildTreeNestsJoinRows(t *testing.T) {
	rows := []treeRow{
		{CompanyID: sp("c1"), CompanyName: sp("Co"), FieldID: sp("f1"), FieldName: sp("F"), SiteID: sp("s1"), SiteName: sp("S"),
			WellID: sp("w1"), WellName: sp("W"), WellboreID: sp("b1"), WellboreName: sp("B"), DesignID: sp("d1"), DesignName: sp("D"),
			TrajectoryID: sp("t1"), TrajectoryName: sp("T"), CaseID: sp("k1"), CaseName: sp("K1")},
		{CompanyID: sp("c1"), CompanyName: sp("Co"), FieldID: sp("f1"), FieldName: sp("F"), SiteID: sp("s1"), SiteName: sp("S"),
			WellID: sp("w1"), WellName: sp("W"), WellboreID: sp("b1"), WellboreName: sp("B"), DesignID: sp("d1"), DesignName: sp("D"),
			TrajectoryID: sp("t1"), TrajectoryName: sp("T"), CaseID: sp("k2"), CaseName: sp("K2")},
		{CompanyID: sp("c2"), CompanyName: sp("Empty")},
	}
	tree := buildTree(rows)
	if len(tree) != 2 || tree[1].Name != "Empty" || len(tree[1].Children) != 0 {
		t.Fatalf("roots: %+v", tree)
	}
	cases := tree[0].Children[0].Children[0].Children[0].Children[0].Children[0].Children[0].Children
	if len(cases) != 2 || cases[0].Kind != "case" || cases[1].Name != "K2" {
		t.Fatalf("cases: %+v", cases)
	}
}
