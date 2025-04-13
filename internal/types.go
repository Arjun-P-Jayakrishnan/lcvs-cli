package internal

const (
	MetaDirName     = ".lcvs"
	MetaFileName    = "meta.json"
	IgnoreFileName  = ".lcvs-ignore"
	StagingFilePath = "staging.json"
)

type ProjectMeta struct {
	Name          string `json:"name"`
	Author        string `json:"author"`
	Description   string `json:"description"`
	CreatedAt     string `json:"created_at"`
	ProjectID     string `json:"project_id"`
	LastKnownPath string `json:"last_known_path"`
}

type StagedFile struct {
	Path string `json:"path"`
	Hash  string `json:"hash"`
}

type StagingInfo struct{
	Files []StagedFile `json:"files"`
}
