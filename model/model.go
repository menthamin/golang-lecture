import (
	"github.com/astaxie/beego/orm"
)

//create database orm_test;
//create table user ( id integer primary key, name varchar(255), profile_id integer );
type User struct {
	Id          int
	Name        string
	// Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
}

// //create table profile ( id integer primary key, age integer, user integer );
// type Profile struct {
// 	Id          int
// 	Age         int16
// 	User        *User   `orm:"reverse(one)"` // Reverse relationship (optional)
// }

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User))
	// orm.RegisterModel(new(User), new(Profile))
}