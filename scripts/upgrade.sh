#!/bin/bash
# Simple script to upgrade the Kubernetes Operator Tutorial to a new version of Kubebuilder.

set -e  # Exit immediately if a command exits with a non-zero status.
set -o pipefail  # Pipefail option to catch errors in pipelines.

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
  echo "Scaffolding the new project..."
  mv operator-v1 operator-v1-old || { echo "Failed to rename existing operator-v1 directory."; exit 1; }
  mkdir operator-v1
  pushd operator-v1 > /dev/null

  kubebuilder init --domain my.domain --repo my.domain/tutorial
  kubebuilder create api --group tutorial --version v1 --kind Foo

  echo "Project scaffolded successfully."
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
scaffold_project

# Print some directives for the user.
cat <<EOF
TODO: Change the <projectName> property to <operator>
TODO: Implement the Foo CRD (<FooSpec> and <FooStatus>)
TODO: Implement the controller (RBAC permissions, reconcile and setupWithManager functions)
NOTE: You may need to resolve some imports such as <corev1>
EOF

wait_for_user_confirmation

# Generate manifests.
echo "Generating manifests..."
if make manifests; then
  echo "Manifests generated successfully."
else
  echo "Failed to generate manifests. Exiting."
  exit 1
fi

# Print other user directives.
cat <<EOF
TODO: Change all occurrences of <operator-v1> to <operator>
TODO: Make sure to keep the operator-v1 title in <README.md>
EOF

popd > /dev/null
echo "Script completed successfully."
