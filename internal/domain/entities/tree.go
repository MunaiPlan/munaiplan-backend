package entities

// TreeNode is one entry of the navigation tree: company → field → site → well → wellbore →
// design → trajectory → case. Kind is the singular entity name used by the frontend routes.
type TreeNode struct {
	ID       string      `json:"id"`
	Kind     string      `json:"kind"`
	Name     string      `json:"name"`
	Children []*TreeNode `json:"children"`
}
