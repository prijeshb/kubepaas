package banner

import (
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/urvil38/kubepaas/config"
)

const deployMessage = 
`
═════════════════════════════════════════════════════════════════════════════════ 
████████████████               Deploying Application             ████████████████
═════════════════════════════════════════════════════════════════════════════════
`

const dockerfileMessage = `
═══════════════════════════════════════════════════════
██████████████ Generating Dockerfile 🐳  ███████████████
═══════════════════════════════════════════════════════
`

const cloudbuildfileMessage = `
═══════════════════════════════════════════════════════
██████████  Generating CloudBuild Config ☁️ 🛠  █████████
═══════════════════════════════════════════════════════
`

const kubernetesMessage = `
═══════════════════════════════════════════════════════
███████████ Generating Kubernetes Config ❅  ███████████
═══════════════════════════════════════════════════════
`

const uploadSourceCode = `
═══════════════════════════════════════════════════
███████████ ⬆︎  Uploading Source Code ⬆︎  ███████████
═══════════════════════════════════════════════════
`

const uploadKubernetes = `
═════════════════════════════════════════════════════════
███████████ ⬆︎  Uploading Kubernetes Config ⬆︎  ███████████
═════════════════════════════════════════════════════════
`

const startCloudLog = `
════════════════════════════════════════ START OF CLOUDBUILD LOG ═════════════════════════════════════════════
`

const endCloudLog = `
════════════════════════════════════════  END OF CLOUDBUILD LOG  ═════════════════════════════════════════════
`

var projectInfo = `
	╔══════════════════╦═══════════════════════════════════════════════╗
	║ Project-Name     ║  %-40s     ║
	╠══════════════════╬═══════════════════════════════════════════════╣
	║ Version          ║  %-40s     ║
	╠══════════════════╬═══════════════════════════════════════════════╣
	║ Runtime          ║  %-40s     ║
	╠══════════════════╬═══════════════════════════════════════════════╣
	║ Source           ║  %-40s     ║
	╠══════════════════╬═══════════════════════════════════════════════╣
	║ URL              ║  %-40s     ║
	╠══════════════════╬═══════════════════════════════════════════════╣
	║ Deployment-Time  ║  %-40s     ║
	╚══════════════════╩═══════════════════════════════════════════════╝
`

func PrintDeployingMessage() string {
	return color.HiGreenString(deployMessage)
}

func PrintDockerfileMessage() string {
	return color.HiCyanString(dockerfileMessage)
}

func SuccessDockerfileMessage() string {
	return color.HiGreenString(" Successfully Generated Dockerfile ✔︎")
}

func PrintCloudBuildMessage() string {
	return color.HiMagentaString(cloudbuildfileMessage)
}

func SuccessDockerCloudbuildMessage() string {
	return color.HiGreenString(" Successfully Generated docker-cloudbuild.json ✔︎")
}

func SuccessKubernetesCloudbuildMessage() string {
	return color.HiGreenString(" Successfully Generated kubernetes-cloudbuild.json ✔︎")
}

func PrintKubernetesMessage() string {
	return color.HiBlueString(kubernetesMessage)
}

func SuccessKubernetesMessage() string {
	return color.HiGreenString(" Successfully Generated kubernetes.yaml ✔︎")
}

func PrintProjectInfo(config config.AppConfig,projectMetadata config.ProjectMetaData) string {
	wd,_:= os.Getwd()
	return color.HiYellowString(
		projectInfo,
		config.ProjectName,
		projectMetadata.CurrentVersion,
		config.Runtime,
		wd,
		"https://"+config.ProjectName+".kubepaas.ml",
		time.Now().Format("2006-01-02 3:4:5 PM"))
}

func StartCloudBuildLog() string {
	return color.HiYellowString(startCloudLog)
}

func EndCloudBuildLog() string {
	return color.HiYellowString(endCloudLog)
}

func PrintUploadSourceCodeMessage() string {
	return color.HiWhiteString(uploadSourceCode)
}

func PrintUploadKubernetesMessage() string {
	return color.HiWhiteString(uploadKubernetes)
}
