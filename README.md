# Fission with TruFaaS

This artifact contains the source code for the modified version of Fission which includes the TruFaaS Internal Component.
You can find the TruFaaS internal component source code in the ``trufaas`` directory.

The rest of this README will guide you through the build and deployment process.

## PreRequisites

Before you can build and run Fission with TruFaaS, you need to have the following software installed on your machine:

1. Go Lang 1.19 or later version (GoLand is recommended as an IDE)
2. Docker (Docker Desktop is recommended if you're using Windows)
3. [Kind (Kubernetes in Docker)](https://kind.sigs.k8s.io/) to create a Kubernetes cluster to run Fission.
4. [Kubectl](https://kubernetes.io/docs/tasks/tools/) and [Helm](https://helm.sh/).
5. [Goreleaser](https://goreleaser.com/install/) to build the Go binaries.
6. [Skaffold](https://skaffold.dev/docs/install/) for a local development workflow to simplify the process of building and deploying Fission.
7. TruFaaS External Component API already running on your local device and [LocalTunnel](https://theboroer.github.io/localtunnel-www/) to a https URL for the API running locally

## Building and Deploying Fission with TruFaaS

1. Before building, you need to specify the URL of the external component inside Fission.
      - Run ```lt --port 8080``` command and keep the terminal open.
      - Copy the https address given
      
   2. Navigate to `trufaas/config.go` and replace ```{https://your-local-tunnel-address}``` in ```ExternalCompBaseURL``` constant with the obtained https address.


2. From back in the main directory run the following commands one after the other:
    ```
    kind create cluster
    kubectl create ns fission
    make skaffold-prebuild # This builds all Go binaries required for Fission
    make create-crds
    skaffold run -p kind
    ```

3. You also need to build the fission-cli separately
   ### Ubuntu
    ```
    GOOS=linux GOARCH=amd64 go build -o fission cmd/fission-cli/main.go
    sudo mv ./fission /usr/local/bin/fission
    ```
   ### Windows
     1. Run ```go build -o fission.exe cmd/fission-cli/main.go```
     2. Move the created `fission.exe` to `C:\Program Files (x86)\fission`.


4. At this point, the modified version of Fission should be up and running. 
To confirm that everything is working properly, run the command ```fission version```. 
If there are no errors returned, then the deployment has been successful.
