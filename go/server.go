package main
import(
  "fmt"
  "net/http"
  "strings"
  "log"
  "time"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "strconv"
  "encoding/json"
)

type User struct
{
  Id_ bson.ObjectId `bson:"_id"`
  Name string `bson:"name"`
  Age int `bson:"age"`
  JoinedAt time.Time `bson:"joned_at"`
  Interests []string `bson:"interests"`
}

func serverCall(res http.ResponseWriter,req *http.Request){
  req.ParseForm() //解析参数，默认是不会解析的  
  fmt.Println(req.Form) //这些信息是输出到服务器端的打印信息  
  fmt.Println("path", req.URL.Path)  
  fmt.Println("scheme", req.URL.Scheme)  
  fmt.Println(req.Form["url_long"])
  user := new(User)  
  for k, v := range req.Form {  
      fmt.Println("key:", k)  
      fmt.Println("val:", strings.Join(v, ""))
      if k == "name" {
        user.Name = v[0]
      } else if k == "age" {
        age,err := strconv.ParseInt(v[0],10,8)
        if err != nil{
           panic(err)
        }
        user.Age = int(age)
      } else if k == "interests"{
        interests := v[0]
        user.Interests = strings.Split(interests,",")
      } else if k == "id_" {
        var id string = v[0]
        user.Id_ = bson.ObjectIdHex(id)
      }
  } 
  
  user.JoinedAt = time.Now()
  path := req.URL.Path
  fmt.Println("req path:" + path)
  if 0 == strings.Compare("/goserver/insert",path) {
    result := insertDb(user)
    fmt.Fprintf(res,result)
  } else if 0 == strings.Compare("/goserver/delete",path) {
    result := deleteById(user.Id_)
    fmt.Fprintf(res,result)
  } else if 0 == strings.Compare("/goserver/update",path) {
    result := updateById(user.Id_)
    fmt.Fprintf(res,result)
  } else if 0 == strings.Compare("/goserver/query",path) {
    user = findById(user.Id_)
    r, _ := json.Marshal(user)
    fmt.Fprintf(res,string(r))
  }
  fmt.Fprintf(res, "Hello go web server") //这个写入到res的是输出到客户端的  
}

//增加数据
func insertDb(user *User) string {
  session,err := mgo.Dial("")
     if err != nil{
     panic(err)
  }
  defer session.Close()
  user.Id_ = bson.NewObjectId()
  c := session.DB("test").C("user")
  err = c.Insert(&user)
  if err != nil{
    panic(err)
  }
  return "插入数据成功"
}
//删除数据
func deleteById(_id bson.ObjectId) string {
  session,err := mgo.Dial("")
  if err != nil{
    panic(err)
  }
  defer session.Close()
  c:= session.DB("test").C("user")
  _,err = c.RemoveAll(bson.M{"_id":_id})
  if err != nil{
    panic(err)
  }
  return "删除数据成功"
}

//更新数据
func updateById(_id bson.ObjectId) string {
  session,err := mgo.Dial("")
  if err != nil{
    panic(err)
  }
  defer session.Close()
  c:= session.DB("test").C("user")
  err = c.Update(bson.M{"_id":_id},bson.M{"$set":bson.M{"age":10}}) 
  if err != nil{
    panic(err)
  } 
  return "更新数据成功"
}

func findById(_id bson.ObjectId) *User {
  session,err := mgo.Dial("")
  if err != nil{
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic,true)
  c := session.DB("test").C("user")
  user := new (User)
  err = c.FindId(_id).One(&user)
  if err != nil{
    panic(err)
  }
  return user
} 

func main(){
  http.HandleFunc("/", serverCall) //设置访问的路由  
  err := http.ListenAndServe(":9090", nil) //设置监听的端口  
  if err != nil {  
      log.Fatal("ListenAndServe: ", err)  
  }
}

