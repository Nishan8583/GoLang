package main

import (
      "fmt";
      bencode "github.com/jackpal/bencode-go"
      "os";
      "log";
      "strconv";
      "crypto/sha1";
      "bytes";
      "time";
      "net/url";
      "net/http";
      "strings";
)
// File available as part of the torrent
type File struct {
    Length int    `bencode:"length"`
    Md5sum string `bencode:"md5sum"`
    Path   string `bencode:"path"`
}

// Data about the download itself
type MetaInfoData struct {
    Name        string `bencode:"name"`
    PieceLength int    `bencode:"piece length"`
    Pieces      string `bencode:"pieces"`
    Private     int    `bencode:"private"`
    Length      int    `bencode:"length"`
    Md5sum      string `bencode:"md5sum"`
    Files       []File `bencode:"files"`
}

type metainfo struct {
      Announce     string       `bencode:"announce"`
      AnnounceList [][]string   `bencode:"announce-list"`
      Info         MetaInfoData `bencode:"info"`
      Encoding     string       `bencode:"encoding"`
      CreationDate int          `bencode:"creation date"`
      CreatedBy    string       `bencode:"created by"`
}

var info metainfo;
// torrent_data() decodes the file and gets info_hash
func torrent_data(filename string) (map[string]interface{}) {

      file, err := os.Open(filename);
      if err != nil {
            log.Fatal("Error Could not r}ead the file, no point in existing ",err);
      }
      defer file.Close();

      data, err := bencode.Decode(file);
      if err != nil {
            log.Fatal("Error occured while decoding the file ",err);
      }

      return data.(map[string]interface{})
}

func get_metadata(filename string) {
      file, err := os.Open(filename);
      if err != nil {
            log.Fatal("Error Could not r}ead the file, no point in existing ",err);
      }
      defer file.Close();

      _ = bencode.Unmarshal(file,&info);
}
func tracker_ready(data map[string]interface{}) {
      announce_url := data["announce"].(string);

      // Used some of the code from https://halfbyte.io/a-bittorrent-client-in-go-part-1-torrent-file-and-announcement/
      infoBuffer := bytes.Buffer{};
      err := bencode.Marshal(&infoBuffer,data["info"]);
      if err != nil {
            fmt.Println("Could not ",err);
      }
      h := sha1.New();
      h.Write(infoBuffer.Bytes());

      // Url Generation starts here
      base_url,err := url.Parse(announce_url);
      if err != nil {
            fmt.Println("There was an error ,",err);
      }

      values := url.Values{};
      values.Add("info_hash",string(h.Sum(nil)));
      values.Add("peer_id",get_peer_id());
      values.Add("port","6009");
      values.Add("uploaded","0");
      values.Add("downloaded","0");
      values.Add("left",strconv.Itoa(info.Info.Length));
      values.Add("compact","1");
      //fmt.Println(data["info"].assert(map[string]string));
      base_url.RawQuery =values.Encode();
      our := strings.Replace(base_url.String(),"udp","http",1);
      fmt.Println(our);
      resp,err := http.Get(our);
      if err != nil {
            fmt.Println(err);
            return;
      }
      fmt.Println(resp);
}


func get_peer_id() (string) {
      h1 := sha1.New()
      h1.Write([]byte(strconv.FormatInt(time.Now().Unix(),10)+strconv.Itoa(os.Getpid())));
      return string(h1.Sum(nil));
}

// get_url() geerates a url with parameters added
// main() is the main function of this file
func main() {
      torrent_dict := torrent_data("1.torrent");
      get_metadata("1.torrent");
      tracker_ready(torrent_dict);

}
