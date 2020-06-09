package infra

import (
	"io/ioutil"

	"github.com/8treenet/extjson"
	"github.com/8treenet/freedom"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	freedom.Prepare(func(initiator freedom.Initiator) {
		initiator.BindInfra(false, func() *Request {
			return &Request{}
		})
		initiator.InjectController(func(ctx freedom.Context) (com *Request) {
			initiator.GetInfra(ctx, &com)
			return
		})
	})
}

// Request .
type Request struct {
	freedom.Infra
}

// BeginRequest .
func (req *Request) BeginRequest(worker freedom.Worker) {
	req.Infra.BeginRequest(worker)
}

// ReadJSON .
func (req *Request) ReadJSON(obj interface{}) error {
	rawData, err := ioutil.ReadAll(req.Worker.IrisContext().Request().Body)
	if err != nil {
		return err
	}
	err = extjson.Unmarshal(rawData, obj)
	if err != nil {
		return err
	}

	return validate.Struct(obj)
}

// ReadQuery .
func (req *Request) ReadQuery(obj interface{}) error {
	if err := req.Worker.IrisContext().ReadQuery(obj); err != nil {
		return err
	}
	return validate.Struct(obj)
}

// ReadForm .
func (req *Request) ReadForm(obj interface{}) error {
	if err := req.Worker.IrisContext().ReadForm(obj); err != nil {
		return err
	}
	return validate.Struct(obj)
}
