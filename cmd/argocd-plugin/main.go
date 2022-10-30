package main

import (
	"fmt"
	"net/http"

	"github.com/argoproj/argo-cd/v2/pkg/apiclient"

	"github.com/UrielCohen456/argo-workflows-argocd-executor-plugin/internal"
)

func main() {
	client, err := apiclient.NewClient(&apiclient.ClientOptions{
		// TODO: make this configurable by passing a root CA.
		Insecure: true,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to initialize Argo CD API client: %s", err))
	}
	http.HandleFunc("/api/v1/template.execute", argocd.ArgocdPlugin(argocd.NewApiExecutor(client)))
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err.Error())
	}
}