package config

import (
	"net/url"
	"reflect"

	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cast"

	"gitlab.com/distributed_lab/figure"

	"gitlab.com/tokend/keypair"
)

var (
	URLHook = figure.Hooks{
		"url.URL": func(value interface{}) (reflect.Value, error) {
			str, err := cast.ToStringE(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse string")
			}
			u, err := url.Parse(str)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse url")
			}
			return reflect.ValueOf(*u), nil
		},
	}
	KeypairHook = figure.Hooks{
		"keypair.Full": func(value interface{}) (reflect.Value, error) {
			switch v := value.(type) {
			case string:
				kp, err := keypair.ParseSeed(v)
				if err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to parse kp")
				}
				return reflect.ValueOf(kp), nil
			case nil:
				return reflect.ValueOf(nil), nil
			default:
				return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
			}
		},
	}
)
