package main

import (

      "fmt";
      bencode "github.com/jackpal/bencode-go"  // Essentital for bencoding and decoding, used bo torrent files
      "os";
      "log";
      "strconv";
      "crypto/sha1";
      "bytes";
      "time";
      "net"

      //"net/url";
      //"net/http";
      "strings";
      "io/ioutil";


)
func connect_tracker(tracker string) (error) {
      fmt.Println(tracker,"1")
      conn, err := net.Dial("udp",tracker);
      if err != nil {
            fmt.Println("Could not connect to server");
            return err;
      }
      fmt.Println("Connection with Server Estabished");
      defer conn.Close();
      //fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n");
      fmt.Fprintf(conn,"GET /announce?compact=0&downloaded=0&info_hash=%0C%3C%E1%19g%29%C0%A8%AF%C7%90%5Ck%E9%B0%A5%92%7Co%BA&left=0&peer_id=%93%0A%87%EA%BB%DE%C9%D5%3B%92%E3%98j%1Cg%B3%EA%97G-&port=6009&uploaded=0 HTTP/1.1\r\n\r\n")
//udp://163.172.180.68:80/announce?compact=1&downloaded=0&info_hash=%0C%3C%E1%19g%29%C0%A8%AF%C7%90%5Ck%E9%B0%A5%92%7Co%BA&left=0&peer_id=%93%0A%87%EA%BB%DE%C9%D5%3B%92%E3%98j%1Cg%B3%EA%97G-&port=6009&uploaded=0
      // conn.Read() reads a certain amount only, in the below case 1024 bytes
      //content := make([]byte,1024);
      //_,err = conn.Read(content);
      fmt.Println("Sending Data to server")
      content,err := ioutil.ReadAll(conn);
      if err != nil {
            fmt.Println("Could not read from server")
            return err;
      }
      fmt.Println(string(content));
      return nil;
}

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
      AnnounceList [][]string   `bencode:"announce-list"`  // the actual announce we will be using
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
      //announce_url := strings.Replace(data["announce"].(string),"/announce","",1);

      // Used some of the code from https://halfbyte.io/a-bittorrent-client-in-go-part-1-torrent-file-and-announcement/
      infoBuffer := bytes.Buffer{};
      err := bencode.Marshal(&infoBuffer,data["info"]);
      if err != nil {
            fmt.Println("Could not ",err);
      }
      h := sha1.New();
      h.Write(infoBuffer.Bytes());

      // Url Generation starts here
      fmt.Println("Length ",len(info.AnnounceList))
      for _,value := range info.AnnounceList {

            url1 := strings.Replace(value[0],"udp://","",1);
            url1 = strings.Replace(url1,"tcp://","",1)
            url1 = strings.Replace(url1,"http://","",1)
            url1 =strings.Replace(url1,"/announce","",1)
            //fmt.Println(url1)
            err := connect_tracker(url1);
            if err != nil {
                  fmt.Println(err)
            }
      }
/*
      base_url,err := url.Parse(info.AnnounceList[19][0]);
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
      fmt.Println(base_url.String())
      our := strings.Replace(base_url.String(),"udp","http",1);
      //fmt.Println(our);
      //fmt.Println(base_url.String())
      resp,err := http.Get(our);

      fmt.Println(resp.StatusCode)
      body,_ := ioutil.ReadAll(resp.Body);
      fmt.Println(string(body))
      */
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
