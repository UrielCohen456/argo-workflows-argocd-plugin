<div align="center">
  <h1 align="center">Argo CD Executor Plugin</h1>
  <p align="center">An <a href="https://github.com/argoproj/argo-workflows/blob/master/docs/executor_plugins.md">Executor Plugin</a> for <a href="https://argoproj.github.io/argo-workflows/">Argo Workflows</a> that lets you interact with Argo CD servers.</p>
</div>

## Example

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Workflow
metadata:
  generateName: argocd-example-
spec:
  entrypoint: main
  templates:
    - name: main
      plugin:
        argocd:
          actions:
            - - app:
                  sync:
                    apps:
                      - name: guestbook
                      - name: guestbook-backend
```

## Getting Started

### Step 1: Get an Argo CD token

The plugin requires a secret named `argocd-sync-token` with a key called `jwt.txt` containing the Argo CD token. See the [Argo CD documentation](https://argo-cd.readthedocs.io/en/stable/user-guide/projects/#project-roles) for information about generating tokens.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: argocd-sync-token
stringData:
  jwt.txt: <token>
```

After defining the secret, apply it to your cluster:

```shell
kubectl apply -f argocd-sync-token.yaml
```

### Step 2: Install the plugin

```shell
kubectl apply -n argo -f https://raw.githubusercontent.com/crenshaw-dev/argocd-executor-plugin/main/manifests/argocd-executor-plugin-configmap.yaml
```

**Note:** You will have to run the workflow using a service account with appropriate permissions. See [examples/rbac.yaml](examples/rbac.yaml) for an example.

### Step 3: Set your `ARGOCD_SERVER` environment variable

By default, the plugin uses `argocd-server.argocd.svc.cluster.local` for `ARGOCD_SERVER`. If you're using a different
server, you can set the `ARGOCD_SERVER` environment variable in the plugin's configmap.

### Step 4: Run a workflow

```shell
argo submit examples/argocd.yaml --serviceaccount my-service-account --watch
```

## Usage

The `actions` field of the plugin config accepts a nested list of actions. Parent lists are executed sequentially, and 
child lists are executed in parallel. This allows you to run multiple actions in parallel, and multiple groups of 
actions in sequence.

### Running syncs in sequence

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Workflow
metadata:
  generateName: argocd-sequence-example-
spec:
  entrypoint: main
  templates:
    - name: main
      plugin:
        argocd:
          actions:
            - - app:
                  sync:
                    apps:
                      - name: guestbook-backend
            - - app:
                  sync:
                    apps:
                      - name: guestbook-frontend
```

### Setting sync options

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Workflow
metadata:
  generateName: argocd-options-example-
spec:
    entrypoint: main
    templates:
      - name: main
        plugin:
          argocd:
            actions:
              - - app:
                    sync:
                      apps:
                        - name: guestbook-backend
                      options:
                        - ServerSideApply=true
                        - Validate=true
```

### Setting a timeout

Each sync action may be configured with a timeout. The default is no timeout.

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Workflow
metadata:
  generateName: argocd-timeout-example-
spec:
    entrypoint: main
    templates:
      - name: main
        plugin:
          argocd:
            actions:
              - - app:
                    sync:
                      apps:
                        - name: guestbook-backend
                      timeout: 30s
```

## Contributing

Head to the [scripts](CONTRIBUTING.md) directory to find out how to get the project up and running on your local machine for development and testing purposes.
