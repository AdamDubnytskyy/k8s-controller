package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var goBasicCmd = &cobra.Command{
	Use:   "go-basic",
	Short: "Run cobra basic code",
	Run: func(cmd *cobra.Command, args []string) {
		// go basic code to run functions
		k8s := Kubernetes{
			Name:    "kind-dev",
			Version: "v1.33.0",
			Users:   []string{"adam", "miles"},
			NodeNumber: func() int {
				return 4
			},
		}

		// print users
		k8s.GetUsers()

		// add new user to struct
		k8s.AddNewUser("anonymous")

		// print users once again
		k8s.GetUsers()
	},
}

func init() {
	rootCmd.AddCommand(goBasicCmd)
}

type Kubernetes struct {
	Name       string     `json:"name"`
	Version    string     `json:"version"`
	Users      []string   `json:"users,omitempty"`
	NodeNumber func() int `json:"-"`
}

func (k8s Kubernetes) GetUsers() {
	for _, user := range k8s.Users {
		log.Info().Str("get user", user).Msg("ok")
	}
}

func (k8s *Kubernetes) AddNewUser(user string) {
	k8s.Users = append(k8s.Users, user)
}
