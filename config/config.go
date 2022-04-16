package config

var (
	// Prefix for application environment variables
	ENV_PREFIX string = "MOVE_FILES_INTO_DATE_DIRECTORIES"

	// envVarNames is a list that contains all supported Environement Variables to pass to CLI
	envVarNames []string
)

// AddEnvParam : Add an Environment Param
func AddEnvParam(envVarName string) {
	envVarNames = append(envVarNames, envVarName)
}
