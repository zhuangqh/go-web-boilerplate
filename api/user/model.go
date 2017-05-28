package user

import (
    "github.com/asaskevich/govalidator"
    "gopkg.in/mgo.v2/bson"
    "time"
    "github.com/zhuangqh/go-web-boilerplate/db"
)

const ModelName = "user"

type Param struct {
    Name         string          `bson:"name"          json:"name"         valid:"runelength(1|100),required"`
    HeadImgUrl   string          `bson:"head_img_url"  json:"head_img_url" valid:"url"`
}

func (param *Param) Validate() (bool, error) {
    return govalidator.ValidateStruct(param)
}

type Model struct {
    ID           bson.ObjectId   `bson:"_id,omitempty" json:"_id"`
    Param `bson:",inline"`
    Time         time.Time       `bson:"time"          json:"time"`
    Deleted      bool            `bson:"deleted"       json:"-"`
}

func newOne(param *Param) *Model {
    return &Model{
        ID: bson.NewObjectId(),
        Param: *param,
        Time: time.Now(),
        Deleted: false,
    }
}

func GetMany() ([]*Model, *db.OPError) {
    db.MgoSession = db.MgoSession.Clone()
    defer db.MgoSession.Close()

    coll := db.MgoSession.DB(db.Database).C(ModelName)

    var data []*Model
    err := coll.Find(nil).All(&data)

    if err != nil {
        return data, &db.OPError{200, err}
    }

    return data, nil
}

func CreateOne(param *Param) (*Model, *db.OPError) {
    if res, err := param.Validate(); res == false {
        return nil, &db.OPError{400, err}
    }

    obj := newOne(param)

    db.MgoSession = db.MgoSession.Clone()
    defer db.MgoSession.Close()

    coll := db.MgoSession.DB(db.Database).C(ModelName)

    err := coll.Insert(obj)

    if err != nil {
        return nil, &db.OPError{500, err}
    }

    return obj, nil
}
