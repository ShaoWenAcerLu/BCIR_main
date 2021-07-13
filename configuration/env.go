package configuration

import (
	"os"
	"path/filepath"
)

const VERSION = "v1"
const BCIRUrl = "/protection/" + VERSION + "/policies"

const SYSTEM_PORT = 22

var ROOT_DIR, _ = os.Getwd()
var BCIR_ENV_DIR = filepath.Join(ROOT_DIR, "env")
var POLICY_CONFIG = filepath.Join(BCIR_ENV_DIR, "policyConfig.txt")
