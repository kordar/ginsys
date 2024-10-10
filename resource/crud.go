package resource

import (
	"errors"
	"github.com/kordar/gocrud"
	"github.com/kordar/govalidator"
)

type CrudClone interface {
	Target() interface{}
}

func Create(body gocrud.FormBody, db interface{}, model CrudClone, params map[string]string, valid func(v interface{}) error) (interface{}, error) {
	err := body.Unmarshal(model)
	if err != nil {
		return nil, err
	}

	validate := govalidator.GetValidate()
	err = validate.Struct(model)
	if err != nil {
		return nil, err
	}

	target := model.Target()

	if valid != nil {
		err = valid(target)
		if err != nil {
			return nil, err
		}
	}

	exec := gocrud.GetExecute("CREATE", body.DriverName(params), "")
	if exec == nil {
		return nil, errors.New("execution function for 'CREATE' not found")
	}

	if e := exec(db, "", target); e == nil {
		return target, nil
	} else {
		return nil, e.(error)
	}
}

func Updates(body gocrud.FormBody, db interface{}, model CrudClone, params map[string]string, valid func(v interface{}) error) (interface{}, error) {
	err := body.Unmarshal(model)
	if err != nil {
		return nil, err
	}

	// TODO 更新model必须提供有效的更新条件
	db, err = body.QuerySafe(db, params)
	if err != nil {
		return nil, err
	}

	target := model.Target()

	if valid != nil {
		err = valid(target)
		if err != nil {
			return nil, err
		}
	}

	exec := gocrud.GetExecute("UPDATES", body.DriverName(params), "")
	if exec == nil {
		return nil, errors.New("execution function for 'UPDATES' not found")
	}

	if e := exec(db, "", target); e == nil {
		return target, nil
	} else {
		return nil, e.(error)
	}
}
