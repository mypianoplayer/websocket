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

func NewObjectBase(nm string) *ObjectBase {
	currentID++
	return &ObjectBase{
		id:   currentID,
		name: nm,
	}
}

func (o *ObjectBase) Name() string {
	return o.name
}

func GetComponent(o Object, name string) Component {
    
	v := reflect.ValueOf(o)
	f := v.Elem().FieldByName(name)
    if f.CanAddr() {
        return f.Addr().Interface().(Component)
    }
    
    log.Println("canAddr failed")
    return nil
}

func SetupComponent(o Object) {

    for comp := range EachComponent(o) {
        comp.SetObject(o)
        log.Println("set object")
    }
}

func EachComponent(o Object) chan Component {
    
    // _, ok := o.(Object)
    // if !ok {
    //     log.Println("o is not Object")
    //     return nil
    // }
    
    ch := make(chan Component)
    v := reflect.ValueOf(o)
    // log.Println(v.Kind())
    n := v.Elem().NumField()
    i := 0
    go func() {
        for {
            if i >= n {
                close(ch)
                break;
            }

            if v.Elem().Field(i).CanAddr() {
            // log.Println("canaddr  ", v.Elem().Field(i).Type(), v.Elem().Field(i).Kind())

                a := v.Elem().Field(i).Addr()
                if a.CanInterface() {
            // log.Println("caninterface  ", v.Elem().Field(i).Type())
                    in := a.Interface()
                    comp, ok := in.(Component)
                    if ok {
                        log.Println(v.Elem().Field(i).Type(), " OK" )
                        ch <- comp
                    }
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