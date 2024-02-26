package models

type IndexerData struct {
	Name         string  `json:"name"`
	StorageType  string  `json:"storage_type"`
	ShardNum     int     `json:"shard_num"`
	MappingField struct {
		Properties map[string]IndexerPropertyDetail `json:"properties"`
	} `json:"mappings"`
}
