package cview

import (
	"testing"

	"github.com/gdamore/tcell"
)

const (
	treeViewTextA = "Hello, world!"
	treeViewTextB = "Goodnight, moon!"
)

func TestTreeView(t *testing.T) {
	t.Parallel()

	tr, sc, err := testTreeView()
	if err != nil {
		t.Error(err)
	}
	if tr.GetRoot() != nil {
		t.Errorf("failed to initalize TreeView: expected nil root node, got %v", tr.GetRoot())
	} else if tr.GetCurrentNode() != nil {
		t.Errorf("failed to initalize TreeView: expected nil current node, got %v", tr.GetCurrentNode())
	} else if tr.GetRowCount() != 0 {
		t.Errorf("failed to initalize TreeView: incorrect row count: expected 0, got %d", tr.GetRowCount())
	}

	rootNode := NewTreeNode(treeViewTextA)
	if rootNode.GetText() != treeViewTextA {
		t.Errorf("failed to update TreeView: incorrect node text: expected %s, got %s", treeViewTextA, rootNode.GetText())
	}

	tr.SetRoot(rootNode)
	tr.Draw(sc)
	if tr.GetRoot() != rootNode {
		t.Errorf("failed to initalize TreeView: expected root node A, got %v", tr.GetRoot())
	} else if tr.GetRowCount() != 1 {
		t.Errorf("failed to initalize TreeView: incorrect row count: expected 1, got %d", tr.GetRowCount())
	}

	tr.SetCurrentNode(rootNode)
	if tr.GetCurrentNode() != rootNode {
		t.Errorf("failed to initalize TreeView: expected current node A, got %v", tr.GetCurrentNode())
	}

	childNode := NewTreeNode(treeViewTextB)
	if childNode.GetText() != treeViewTextB {
		t.Errorf("failed to update TreeView: incorrect node text: expected %s, got %s", treeViewTextB, childNode.GetText())
	}

	rootNode.AddChild(childNode)
	tr.Draw(sc)
	if tr.GetRoot() != rootNode {
		t.Errorf("failed to initalize TreeView: expected root node A, got %v", tr.GetRoot())
	} else if tr.GetRowCount() != 2 {
		t.Errorf("failed to initalize TreeView: incorrect row count: expected 1, got %d", tr.GetRowCount())
	}
}

func testTreeView() (*TreeView, tcell.Screen, error) {
	t := NewTreeView()

	sc := tcell.NewSimulationScreen("UTF-8")
	sc.SetSize(80, 24)

	_ = NewApplication().SetRoot(t, true).SetScreen(sc)

	return t, sc, nil
}
