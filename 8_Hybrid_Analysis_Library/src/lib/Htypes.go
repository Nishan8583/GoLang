// Here lies all the types needed for json unmarshalling
package Hybrid;

// Json unmarshalling type for /search/hash
// Api Reference: https://www.hybrid-analysis.com/docs/api/v2#/Search/post_search_hash
type mitre struct {
    Tactic string `json:"tactic"`
    Technique string `json:technique`
    AttckId string `json:"attck_id"`
    AttckIdWiki string `json:"attck_id_wiki"`
    MaliciouIDentifiersCount int `json:"malicious_identifiers_count"`
    MaliciousIdentifiers []string `json:"malicious_identifiers"`
    SuspiciousIdentifiers []string `json:"suspicious_identifiers"`
    SuspiciousIdentifiersCount int `json:"suspicious_identifiers_count"`
    InformativeIdentifiers []string `json:"informative_identifiers"`
    InformativeIdentifiersCount int `json:informative_identifiers_count`
}
type SearchHashType struct {
    JobId string `json:"job_id"`
    EnviromentId int `json:"environment_id"`
    Environment_description string `json:"environment_description"`
    Size int `json:"size"`
    Stype string `json:"type"`
    StypeShort []string `json:"type_short"`
    TargetURL string `json:"target_url"`
    State string `json:"state"`
    ErrorType string `json:"error_type"`
    ErrorOrigin string `json:"error_origin"`
    SubmitName string `json:"submit_name"`
    Md5 string `json:"md5"`
    Sha1 string `json:"sha1"`
    Sha256 string `json:"sha256"`
    Sha512 string `json:"sha512"`
    Ssdeep string `json:"ssdeep"`
    Imphash string `json:"imphash"`
    AvDetect int `json:"av_detect"`
    Vxfamily string `json:"vx_family"`
    UrlAnalysis bool `json:"url_analysis"`
    Analysis_start_time string `json:"analysis_start_time"`
    ThreatScore int `json:"threat_score"`
    Interesting bool `json:interesting`
    ThreatLevel int `json:"threat_level"`
    Verdict string `json:"verdict"`
    Certificates []string `json:"certificates"`
    Domains []string `json:domains`
    ClassificationTags []string `json:"classification_tags"`
    CompromisedHosts []string `json:"compromised_hosts"`
    Hosts []string `json:"hosts"`
    TotalNetworkConnections int `json:"total_network_connections"`
    TotalProcesses int `json:"total_processes"`
    TotalSignatures int `json:"total_signatures"`
    Processes []string `json:"processes"`
    FileMetadata []string `json:"file_metadata"`
    Tags []string `json:"tags"`
    MitreAttcks []mitre `json:"mitre_attcks"`

}
