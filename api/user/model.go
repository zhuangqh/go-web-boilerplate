package user

import (
    "github.com/asaskevich/govalidator"
    . "github.com/zhuangqh/go-web-boilerplate/db"
)

const ModelName = "user"

type Param struct {
    Name         string          `json:"name"         valid:"runelength(1|100),required"`
    HeadImgUrl   string          `json:"head_img_url" valid:"url"`
}

func (param *Param) Validate() (bool, error) {
    return govalidator.ValidateStruct(param)
}

type Model struct {
    BasicModel
    Param
}

func (Model) TableName() string {
    return ModelName
}

func init() {
    DB.AutoMigrate(&Model{})
}

/**
 * methods
 */

func GetMany() ([]*Model, *OPError) {
    var objs []*Model

    DB.Find(&objs)

    return objs, nil
}

func CreateOne(param *Param) (*Model, *OPError) {
    if res, err := param.Validate(); !res {
        return nil, &OPError{400, err}
    }

    obj := Model{
        Param: *param,
    }

    err := DB.Create(&obj).Error

    if err != nil {
        return nil, &OPError{500, err}
    }

    return &obj, nil
}
