// Here lies all the types needed for json unmarshalling
package Hybrid;

//The base type the package will be using which will contain the httpRequest state and client to do interaction
import (
  "net/http";
)
type GoHybrid struct {
  req *http.Request;
  client http.Client;
}

// Json unmarshalling type for /search/hash
// Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Search/post_search_hash
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
    Processes []map[string]string `json:"processes"` // An error may occur since if no value null is returned
    FileMetadata []string `json:"file_metadata"`
    Tags []string `json:"tags"`
    MitreAttcks []mitre `json:"mitre_attcks"`

}

// Json Unmarshalling for /search/SearchTerms
// Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Search/post_search_terms*/
type SearchTermsResult struct {
  Verdict string `json:"verdict"`
  Vxfamily string `json:"vx_family"`
  Sha256 string `json:"sha256"`
  ThreatScore int `json:"threat_score"`
  JobId string `json:"job_id"`
  EnviromentId int `json:"environment_id"`
  Analysis_start_time string `json:"analysis_start_time"`
  SubmitName string `json:"submit_name"`
  Environment_description string `json:"environment_description"`
  Size int `json:"size"`
  Type string `json:type`
  TypeShort string `json:type_short`
}
type SearchTermsType struct {
  SearchTerms []map[string]string `json:"search_terms"`
  Count int `json:"count"`
  Result []SearchTermsResult `json:"result"`

}

// Json Unmarshalling type for /overview/{sha256}
// Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Analysis_Overview/get_overview__sha256_
type OverviewTypeScanner struct {
  Name string `json:"name"`
  Status string `json:"status"`
  Progress int `json:"progress"`
  Total int `json:"total"`
  Positives int `json:"positives"`
  Percent int `json:"percent"`
  AntiVirusResults []string `json:"anti_virus_results"`
}
type OverviewType struct {
  Sha256 string `json:"sha256"`
  LastFileName string `json:"last_file_name"`
  OtherFileName []string `json:"other_file_name"`
  ThreatScore int `json:"threat_score"`
  Verdict string `json:"verdict"`
  UrlAnalysis bool `json:"url_analysis"`
  Size int `json:"size"`
  Type string `json:"type"`
  TypeShort []string `json:"type_short"`
  AnalysisStartTime string `json:"analysis_start_time"`
  LastMultiScan string `json:"last_multi_scan"`
  Tags []string `json:"tags"`
  Architecture string `json:"architecture"`
  MultiScanResult int `json:"multiscan_result"`
  Scanners []OverviewTypeScanner `json:"scanners"`
  RelatedParentHashes []string `json:"related_parent_hashes"`
  RelatedChilrenHahses []string `json:"related_children_hashes"`
  WhiteListed bool `json:"whitelisted"`
  ChildrenInQueue int `json:"children_in_queue"`
  Children_in_progress int `json:"children_in_progress"`
  RelatedReports []string `json:"related_reports"`
}

// Json Unmarshalling for overview/{sha256}/summary
// Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Analysis_Overview/get_overview__sha256__summary
type OverviewSummaryType struct {
    Sha256 string `json:"sha256"`
    ThreatScore int `json:"threat_score"`
    Verdict string `json:"verdict"`
    AnalysisStartTime string `json:"analysis_start_time"`
    LastMultiScan string `json:"last_multi_scan"`
    MultiScanResult int `json:"multiscan_result"`
}

// Json unmarshalling for /report/{id}/state
// Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Report/get_report__id__state
type ReportStateType struct {
    State string `json:"state"`
    ErrorType string `json:"error_type"`
    ErrorOrigin string `json:"error_origin"`
    Error string `json:"error"`
    RelatedReports []string `json:"related_reports"`
}

// Json Unmarshalling for /report/{id}/summary
// Reference Api: https://www.hybrid-analysis.com/docs/api/v2#/Sandbox_Report/get_report__id__summary
type ExtractedFiles struct {
    Name string `json:"name"`
    FilePath string `json:"file_path"`
    Sha1 string `json:"sha1"`
    Sha256 string `json:"sha256"`
    Md5 string `json:"md5"`
    TypeTags []string `json:"type_tags"`
    Description string `json:"description"`
    RuntimeProcess string `json:"runtime_process"`
    ThreatLevel int `json:"threat_level"`
    ThreatLevelReadable string `json:"threat_level_readable"`
    AvLabel string `json:"av_label"`
    AvMatched int `json:"av_matched"`
    AvTotal int `json:"av_total"`
    FileAvailableToDownload bool `json:"file_available_to_download"`
}
type ReportSummaryIDType struct {
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
    Extracted_Files []ExtractedFiles `json:"extracted_files"`
    Processes []map[string]string `json:"processes"`  // An error may occur since if no value null is returned
    FileMetadata []string `json:"file_metadata"`
    Tags []string `json:"tags"`
    MitreAttcks []mitre `json:"mitre_attcks"`

}
