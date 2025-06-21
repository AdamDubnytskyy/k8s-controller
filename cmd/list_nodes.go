package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Status string
type Role string

type Node struct {
	Name    string `json:"name"`
	Status  Status `json:"status"`
	Role    Role   `json:"role"`
	Age     string `json:"age"`
	Version string `json:"version"`
}

const (
	StatusReady    Status = "Ready"
	StatusNotReady Status = "NotReady"
)

const (
	RoleControlPlane Role = "control-plane"
	RoleDataPlane    Role = "data-plane"
)

var fetchListOfNodesCmd = &cobra.Command{
	Use:   "list-nodes",
	Short: "Fetch list of nodes",
	Run: func(cmd *cobra.Command, args []string) {
		node := Node{
			Name:    "dev-control-plane",
			Status:  StatusReady,
			Role:    RoleControlPlane,
			Age:     "1d",
			Version: "v1.33.0",
		}

		// fetch list of cluster nodes
		node.GetAllNodes()

		// fetch cluster control-plane nodes
		node.GetControlPlaneNodes()

		// fetch cluster data-plane nodes
		node.GetDataPlaneNodes()
	},
}

func init() {
	rootCmd.AddCommand(fetchListOfNodesCmd)
}

func (node *Node) GetAllNodes() {
	fmt.Println("GetAllNodes")
}

func (node *Node) GetControlPlaneNodes() {
	fmt.Println("GetControlPlaneNodes")
}

func (node *Node) GetDataPlaneNodes() {
	fmt.Println("GetDataPlaneNodes")
}
