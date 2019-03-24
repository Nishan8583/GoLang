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
)

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

func tracker_ready(data map[string]interface{}) {
      announce_url := data["announce"];
      infoBuffer := bytes.Buffer{};
      err := bencode.Marshal(&infoBuffer,data["info"]);
      if err != nil {
            fmt.Println("Could not ",err);
      }
      h := sha1.New();
      h.Write(infoBuffer.Bytes());
      fmt.Println(h.Sum(nil),announce_url);
}

func get_peer_id() {
      h1 := sha1.New()
      h1.Write([]byte(strconv.FormatInt(time.Now().Unix(),10)+strconv.Itoa(os.Getpid())));
      fmt.Println(h1.Sum(nil));
}
// main() is the main function of this file
func main() {
      torrent_dict := torrent_data("1.torrent");
      tracker_ready(torrent_dict);
      get_peer_id();

}
