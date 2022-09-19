package config

type ContainerInfo struct {
	ContainerName string
	ContainerId   string
}

// singleton pattern implemented
var containerInfo *ContainerInfo

func GetContainerInfo() *ContainerInfo {
	if containerInfo == nil {
		containerInfo = &ContainerInfo{}
	}
	return containerInfo
}
