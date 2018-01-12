package main
import(
  "fmt"
  "time"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type User struct
{
  Id_ bson.ObjectId `bson:"_id"`
  Name string `bson:"name"`
  Age int `bson:"age"`
  JoinedAt time.Time `bson:"joned_at"`
  Interests []string `bson:"interests"`
}
func PrintUser(users []User){
  for _,v:= range users{
    fmt.Println("Name:",v.Name,"Age:",v.Age,"Interests:",v.Interests)
  }
}
func main(){
    session,err := mgo.Dial("")
    if err != nil{
      panic(err) 
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic,true)

    c := session.DB("test").C("user")
    
    var users []User 
    c.Find(nil).All(&users)
    /*for _,v:= range users{
      fmt.Println("Name:",v.Name,"Age:",v.Age,"Interests:",v.Interests)
    }*/
    //fmt.Println(users)
    PrintUser(users)
    fmt.Println("------------------------------------------------------")
    id := "5a3c9c37e13823393b37b8b6"
    objectId := bson.ObjectIdHex(id)
    fmt.Println(">>>>>>>>>>>>>>>>>>>>>>:"+objectId)
    user := new (User)
    err =  c.Find(bson.M{"_id":objectId}).One(&user)
    if err != nil{
      panic(err)
    }
    fmt.Println(user)
    fmt.Println("-------------------------------------------------------- find by id")
    objectId = bson.ObjectIdHex("5a3c9c37e13823393b37b8b6")
    err = c.FindId(objectId).One(&user)
    if err != nil{
      panic(err)
    }
    fmt.Println(user)
    fmt.Println("-------------------------------------------------------- find by condition $eq")
    err = c.Find(bson.M{"name":"tony"}).All(&users)
    if err != nil{
      panic(err)
    }
    //fmt.Println(users)
    PrintUser(users)
    fmt.Println("--------------------------------------------------------find by condition $ne")
    err = c.Find(bson.M{"name":bson.M{"$ne":"tony"}}).All(&users)
    if err != nil {
      panic(err)
    }
    PrintUser(users)
    //fmt.Println(users)
    fmt.Println("-------------------------------------------------------find by condition $gt")
    err = c.Find(bson.M{"age":bson.M{"$gt":28}}).All(&users)
    if err != nil{
      panic(err)
    }
    //fmt.Println(users)
    PrintUser(users)
    fmt.Println("-------------------------------------------------------find by condition $lt")
    err = c.Find(bson.M{"age":bson.M{"$lt":28}}).All(&users)
    if err != nil{
      panic(err)
    }
    //fmt.Println(users)
    PrintUser(users)
    fmt.Println("-------------------------------------------------------find by condition $gte")
    err = c.Find(bson.M{"age":bson.M{"$gte":28}}).All(&users)
    if err != nil{
      panic(err)
    }
    //fmt.Println(users)
    PrintUser(users)
    fmt.Println("------------------------------------------------------find by condition $lte")
    err = c.Find(bson.M{"age":bson.M{"$lte":28}}).All(&users)
    if err != nil{
      panic(err)
    }
    //fmt.Println(users)
    PrintUser(users)
    fmt.Println("------------------------------------------------------find by condition $in")
    err = c.Find(bson.M{"name":bson.M{"$in":[]string{"lisi","tony"}}}).All(&users)
    if err != nil{
      panic(err)
    }
    //fmt.Println(users)
    PrintUser(users)
    fmt.Println("-----------------------------------------------------find by mutl condition")
    err = c.Find(bson.M{"name":"tony","age":45}).All(&users)
    if err != nil{
      panic(err)
    }
    //fmt.Println(users)
    PrintUser(users)
    fmt.Println("----------------------------------------------------find by mutl condition or ")
    err = c.Find(bson.M{"$or":[]bson.M{bson.M{"name":"tony"},bson.M{"age":15}}}).All(&users)
    if err != nil{
      panic(err)
    }
    //fmt.Println(users)
    PrintUser(users)
}


