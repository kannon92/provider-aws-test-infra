package config

import "embed"

//go:embed ubuntu configure.sh run-kubeadm.sh run-post-install.sh al2023.sh *.yaml
var ConfigFS embed.FS
