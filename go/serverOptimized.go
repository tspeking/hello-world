package main
import(
  "fmt"
  "net/http"
  "strings"
  "log"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "encoding/json"
)

func serverCall(res http.ResponseWriter,req *http.Request){
  req.ParseForm() //解析参数，默认是不会解析的  
  fmt.Println(req.Form) //这些信息是输出到服务器端的打印信息  
  fmt.Println("path", req.URL.Path)  
  fmt.Println("scheme", req.URL.Scheme)  
  fmt.Println(req.Form["url_long"])
  
  path := req.URL.Path
  if 0 == strings.Compare("/goserver/add",path) {
    result := addPerson(req)
    fmt.Fprintf(res,result)
  } else if 0 == strings.Compare("/goserver/delete",path) {
    result := delPerson(req)
    fmt.Fprintf(res,result)
  } else if 0 == strings.Compare("/goserver/update",path) {
    result := updatePerson(req)
    fmt.Fprintf(res,result)
  } else if 0 == strings.Compare("/goserver/query",path) {
    result := queryPerson(req)
    fmt.Fprintf(res,result)
  } else if 0 == strings.Compare("/goserver/queryPersonMulti",path) {
    result := queryPersonMulti(req)
    fmt.Fprintf(res,result)
  }

  fmt.Fprintf(res, "Hello go web server")
}

type Person struct {
  Id bson.ObjectId `bson:"_id"`
  Name string `bson:"name"`
  Phone string `son:"phone"` 
}

var mgoSession *mgo.Session
var dbBase = "test"

/**
*获取session,如果存在则拷贝一份
*/
func getSession() *mgo.Session {
  if mgoSession == nil {
    var err error
    mgoSession,err = mgo.Dial("")
    if err != nil {
      panic(err)
    }
  }
  //默认最大连接池为4096
  return mgoSession.Clone()
}

func doOperation(collection string,excute func(*mgo.Collection) error) error {
  session := getSession()
  defer session.Close()
  c := session.DB(dbBase).C(collection)
  return excute(c)
}

//插入数据
func addPerson(req *http.Request) string {
  var person = new(Person)
  person.Id = bson.NewObjectId()
  for k, v := range req.Form {
    fmt.Println("key:", k,"value:",strings.Join(v, ""))
    if k == "name" {
        person.Name = v[0]
    } else if k == "phone" {
        person.Phone = v[0]
    } 
  }
  query := func(c *mgo.Collection) error {
    return c.Insert(person)
  }

  err := doOperation("person",query)
  if err != nil {
    fmt.Println(err)
    return "add Person fail! "
  }
  return "add Person success!  "
}

//删除数据
func delPerson(req *http.Request) string {
  var id string
  for k,v := range req.Form {
    if k == "id_" {
      id = v[0]
      break
    }
  }
  query := func(c *mgo.Collection) error {
    _,err := c.RemoveAll(bson.M{"_id":bson.ObjectIdHex(id)})
    return err
  }
  err := doOperation("person",query)
  if err != nil {
    fmt.Println(err)
    return "del Person fail! "
  } else {
    return "del Person success! "
  }
}

//更新数据
func updatePerson(req *http.Request) string {
  var id string
  var phone string
  var name string
  for k,v := range req.Form {
    if k == "id_" {
      id = v[0]
    } else if k == "phone" {
      phone = v[0]
    } else if k == "name" {
      name = v[0]
    }
  }
  query := func(c *mgo.Collection) error {
    return c.Update(bson.M{"_id":bson.ObjectIdHex(id),"name":name},bson.M{"$set":bson.M{"phone":phone}})
  }
  err := doOperation("person",query)
  if err != nil {
    fmt.Println(err)
    return "update Person fail!  "
  } else {
    return "update Person success!  "
  }
}

//查询数据
func queryPerson(req *http.Request) string {
  var id string
  for k,v := range req.Form {
    if k == "id_" {
      id = v[0]
      break
    }
  }
  var person Person
  query := func(c *mgo.Collection) error {
    return c.FindId(bson.ObjectIdHex(id)).One(&person)
  }
  err := doOperation("person",query)
  if err != nil {
    fmt.Println(err)
    return "query Person fail! "
  }
  r, _ := json.Marshal(person)
  return string(r)
}

//多条件查询
func queryPersonMulti(req *http.Request) string {
  var name string
  var phone string
  for k,v := range req.Form {
    if k == "name" {
      name = v[0]
    } else if k == "phone" {
      phone = v[0]
    }
  }
  var persons []Person
  query := func(c *mgo.Collection) error {
    return c.Find(bson.M{"name":name,"phone":phone}).All(&persons)
  }
  err := doOperation("person",query)
  if err != nil {
    fmt.Println(err)
    return "query Persons fail! "
  } else {
    r,_:=json.Marshal(persons)
    return string(r)
  }
}

func main() {
  http.HandleFunc("/", serverCall) //设置访问的路由  
  err := http.ListenAndServe(":9090", nil) //设置监听的端口  
  if err != nil {  
      log.Fatal("ListenAndServe: ", err)  
  }
}
