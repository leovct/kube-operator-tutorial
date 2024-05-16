#!/bin/bash

# Simple script to upgrade the Kubernetes Operator Tutorial to a new version of Kubebuilder.

# Exit immediately if a command exits with a non-zero status.
set -e
# Pipefail option to catch errors in pipelines.
set -o pipefail

# Function to download and install the latest version of Kubebuilder.
install_kubebuilder() {
  echo "Downloading the latest version of Kubebuilder..."
  if curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)"; then
    echo "Download complete."
    chmod +x kubebuilder
    sudo mv kubebuilder /usr/local/bin
    echo "Kubebuilder installed successfully."
  else
    echo "Failed to download Kubebuilder. Exiting."
    exit 1
  fi
}

# Function to scaffold the new project.
scaffold_project() {
  local name=$1

  echo "Scaffolding the new project $name..."
  mv "$name" "${name}-old" || { echo "Failed to rename existing $name directory."; exit 1; }
  mkdir "$name"
  pushd "$name" > /dev/null || exit

  kubebuilder init --domain my.domain --repo my.domain/tutorial
  kubebuilder create api --group tutorial --version v1 --kind Foo

  find . -type f \( -name "*.go" -o -name "*.yaml" -o -name "*.md" -o -name "README.md" -o -name "PROJECT" \) -exec sed -i '' "s/$name/operator/g" {} +
  sed -i '' "s/# operator/# $name/g" README.md
  cp "../${name}-old/config/samples/tutorial_v1_foo.yaml" config/samples
  cp "../${name}-old/config/samples/pod.yaml" config/samples

  echo "Project $name scaffolded successfully."
}

# Function to wait for user confirmation.
wait_for_user_confirmation() {
  while true; do
    read -p "Are you done? (y/n): " response
    case "$response" in
      y|Y)
        echo "Great! Let's carry on!"
        break
        ;;
      n|N)
        echo "Please respond with 'y' to proceed."
        ;;
      *)
        echo "Invalid input. Please respond with 'y' or 'n'."
        ;;
    esac
  done
}

# Main script execution.
install_kubebuilder

if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <project_name>"
  exit 1
fi
project_name=$1
scaffold_project "$project_name"

# Print some directives for the user.
cat <<EOF
TODO: Implement the Foo CRD (<FooSpec> and <FooStatus>)
TODO: Implement the controller (RBAC permissions, reconcile and setupWithManager functions)
TODO: Resolve imports such as <corev1>
EOF

wait_for_user_confirmation

echo "Generating manifests..."
make manifests

echo "Building binary..."
make build

popd > /dev/null
echo "Script completed successfully."
