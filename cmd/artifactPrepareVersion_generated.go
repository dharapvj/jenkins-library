// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/piperenv"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type artifactPrepareVersionOptions struct {
	BuildTool           string `json:"buildTool,omitempty"`
	DockerVersionSource string `json:"dockerVersionSource,omitempty"`
	FilePath            string `json:"filePath,omitempty"`
	GitUserEMail        string `json:"gitUserEMail,omitempty"`
	GitUserName         string `json:"gitUserName,omitempty"`
	IncludeCommitID     bool   `json:"includeCommitId,omitempty"`
	Password            string `json:"password,omitempty"`
	TagPrefix           string `json:"tagPrefix,omitempty"`
	Username            string `json:"username,omitempty"`
	VersioningTemplate  string `json:"versioningTemplate,omitempty"`
	VersioningType      string `json:"versioningType,omitempty"`
}

type artifactPrepareVersionCommonPipelineEnvironment struct {
	artifactVersion string
}

func (p *artifactPrepareVersionCommonPipelineEnvironment) persist(path, resourceName string) {
	content := []struct {
		category string
		name     string
		value    string
	}{
		{category: "", name: "artifactVersion", value: p.artifactVersion},
	}

	errCount := 0
	for _, param := range content {
		err := piperenv.SetResourceParameter(path, resourceName, filepath.Join(param.category, param.name), param.value)
		if err != nil {
			log.Entry().WithError(err).Error("Error persisting piper environment.")
			errCount++
		}
	}
	if errCount > 0 {
		os.Exit(1)
	}
}

// ArtifactPrepareVersionCommand Prepares and potentially updates the artifact's version before building the artifact.
func ArtifactPrepareVersionCommand() *cobra.Command {
	metadata := artifactPrepareVersionMetadata()
	var stepConfig artifactPrepareVersionOptions
	var startTime time.Time
	var commonPipelineEnvironment artifactPrepareVersionCommonPipelineEnvironment

	var createArtifactPrepareVersionCmd = &cobra.Command{
		Use:   "artifactPrepareVersion",
		Short: "Prepares and potentially updates the artifact's version before building the artifact.",
		Long: `Prepares and potentially updates the artifact's version before building the artifact.

The continuous delivery process requires that each build is done with a unique version number.

The version generated using this step will contain:

* Version (major.minor.patch) from descriptor file in master repository is preserved. Developers should be able to autonomously decide on increasing either part of this version number.
* Timestamp
* CommitId (by default the long version of the hash)

Optionally, but enabled by default, the new version is pushed as a new tag into the source code repository (e.g. GitHub).
If this option is chosen, git credentials and the repository URL needs to be provided.
Since you might not want to configure the git credentials in Jenkins, committing and pushing can be disabled using the ` + "`" + `commitVersion` + "`" + ` parameter as described below.
If you require strict reproducibility of your builds, this should be used.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			startTime = time.Now()
			log.SetStepName("artifactPrepareVersion")
			log.SetVerbose(GeneralConfig.Verbose)
			return PrepareConfig(cmd, &metadata, "artifactPrepareVersion", &stepConfig, config.OpenPiperFile)
		},
		Run: func(cmd *cobra.Command, args []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				commonPipelineEnvironment.persist(GeneralConfig.EnvRootPath, "commonPipelineEnvironment")
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, "artifactPrepareVersion")
			artifactPrepareVersion(stepConfig, &telemetryData, &commonPipelineEnvironment)
			telemetryData.ErrorCode = "0"
		},
	}

	addArtifactPrepareVersionFlags(createArtifactPrepareVersionCmd, &stepConfig)
	return createArtifactPrepareVersionCmd
}

func addArtifactPrepareVersionFlags(cmd *cobra.Command, stepConfig *artifactPrepareVersionOptions) {
	cmd.Flags().StringVar(&stepConfig.BuildTool, "buildTool", os.Getenv("PIPER_buildTool"), "Defines the tool which is used for building the artifact.")
	cmd.Flags().StringVar(&stepConfig.DockerVersionSource, "dockerVersionSource", os.Getenv("PIPER_dockerVersionSource"), "For Docker only: Specifies the source to be used for for generating the automatic version. * This can either be the version of the base image - as retrieved from the `FROM` statement within the Dockerfile, e.g. `FROM jenkins:2.46.2` * Alternatively the name of an environment variable defined in the Docker image can be used which contains the version number, e.g. `ENV MY_VERSION 1.2.3`")
	cmd.Flags().StringVar(&stepConfig.FilePath, "filePath", os.Getenv("PIPER_filePath"), "Defines a custom path to the descriptor file. Build tool specific defaults are used (e.g. maven: pom.xml, npm: package.json, mta: mta.yaml)")
	cmd.Flags().StringVar(&stepConfig.GitUserEMail, "gitUserEMail", os.Getenv("PIPER_gitUserEMail"), "Allows to overwrite the global git setting 'user.email' available on your Jenkins server.")
	cmd.Flags().StringVar(&stepConfig.GitUserName, "gitUserName", os.Getenv("PIPER_gitUserName"), "Allows to overwrite the global git setting 'user.name' available on your Jenkins server.")
	cmd.Flags().BoolVar(&stepConfig.IncludeCommitID, "includeCommitId", true, "Defines if the automatically generated version (versioningType 'cloud') should include the commit id hash .")
	cmd.Flags().StringVar(&stepConfig.Password, "password", os.Getenv("PIPER_password"), "Password/token for git authentication")
	cmd.Flags().StringVar(&stepConfig.TagPrefix, "tagPrefix", "build_", "Defines the prefix which is used for the git tag which is written during the versioning run.")
	cmd.Flags().StringVar(&stepConfig.Username, "username", os.Getenv("PIPER_username"), "User name for git authentication")
	cmd.Flags().StringVar(&stepConfig.VersioningTemplate, "versioningTemplate", os.Getenv("PIPER_versioningTemplate"), "DEPRECATED: Defines the template for the automatic version which will be created")
	cmd.Flags().StringVar(&stepConfig.VersioningType, "versioningType", "cloud", "Defines the type of versioning (cloud: fully automatic, library: manual, libraryTag: automatic based on latest tag)")

	cmd.MarkFlagRequired("buildTool")
}

// retrieve step metadata
func artifactPrepareVersionMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "artifactPrepareVersion",
			Aliases: []config.Alias{{Name: "artifactSetVersion", Deprecated: false}, {Name: "setVersion", Deprecated: true}},
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "buildTool",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "dockerVersionSource",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "filePath",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "gitUserEMail",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "gitUserName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "includeCommitId",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "password",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "tagPrefix",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "username",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "versioningTemplate",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "versioningType",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
				},
			},
		},
	}
	return theMetaData
}
