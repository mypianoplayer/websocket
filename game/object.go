package game

import (
    "log"
    "reflect"
)

// type ComponentHolder interface {
// 	Component(name string) Component
// 	EachComponent() chan Component
// }

type Object interface {
	Name() string
}

var currentID int

type ObjectBase struct {
	id   int
	name string
}

func NewObject(nm string) *ObjectBase {
	currentID++
	return &ObjectBase{
		id:   currentID,
		name: nm,
	}
}

func (o *ObjectBase) Name() string {
	return o.name
}

func GetComponent(o interface{}, name string) Component {
	v := reflect.ValueOf(o)
	f := v.FieldByName(name)

	return f.Interface().(Component)
}

func SetupComponent(o interface{}) {

    _, ok := o.(Object)
    if !ok {
        return
    }

	v := reflect.ValueOf(o)
	n := v.Elem().NumField()
	for i := 0; i < n; i++ {
		if v.Elem().Field(i).CanAddr() {
			ii := v.Elem().Field(i).Addr().Interface()
			comp, ok := ii.(Component)
			if ok {
				comp.SetObject(o.(Object))
			}
		}
	}
}

func EachComponent(o interface{}) chan Component {
    
    _, ok := o.(Object)
    if !ok {
        return nil
    }
    
    ch := make(chan Component)
    v := reflect.ValueOf(o)
    log.Println(v.Kind())
    n := v.Elem().NumField()
    log.Println("numf", n)
    i := 0
    go func() {
        for {
            if i >= n {
                close(ch)
                break;
            }

            log.Println(v.Elem().Field(i))

            if v.Elem().Field(i).CanAddr() {
                ii := v.Elem().Field(i).Addr().Interface()
                comp, ok := ii.(Component)
                if ok {
                    log.Println("OK")
                    ch <- comp
                }
            }
            i++
        }
    }()

    return ch
}

// func (o *Object) EachComponent() chan Component {
//     ch := make(chan Component)
//     v := reflect.ValueOf(*o)
//     n := v.NumField()
//     log.Println("numf", n)
//     i := 0
//     go func() {
//         for {
//             if i >= n {
//                 close(ch)
//                 break;
//             }

//             ch <- v.Field(i).Interface().(Component)
//             i++
//         }
//     }()

//     return ch
// }