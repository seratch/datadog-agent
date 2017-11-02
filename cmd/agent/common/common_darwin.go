// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2017 Datadog, Inc.

package common

import (
	"path/filepath"
)

const (
	// DefaultConfPath points to the folder containing datadog.yaml
	DefaultConfPath = "/opt/datadog-agent/etc/datadog-agent"
	// DefaultLogFile points to the log file that will be used if not configured
	DefaultLogFile = "/var/log/datadog/agent.log"
)

var (
	// PyChecksPath holds the path to the python checks from integrations-core shipped with the agent
	PyChecksPath = filepath.Join(_here, "..", "..", "checks.d")
	// DistPath holds the path to the folder containing distribution files
	distPath = filepath.Join(_here, "dist")
)

// GetDistPath returns the fully qualified path to the 'dist' directory
func GetDistPath() string {
	return distPath
}

// GetViewsPath returns the fully qualified path to the 'gui/views' directory
func GetViewsPath() string {
	return filepath.Join(distPath, "views")
}
