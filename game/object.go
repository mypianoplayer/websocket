package game

import (
    "log"
    "reflect"
)

type ComponentHolder interface {
	Component(name string) Component
	EachComponent() chan Component
}

type Object struct {
}

func (o *Object) Component(name string) Component {
    v := reflect.ValueOf(o)
    f := v.FieldByName(name)

    return f.Interface().(Component)
}

func EachComponent(o interface{}) chan Component {
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