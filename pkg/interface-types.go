// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package pkg

import "reflect"

// ReflectTypeIHello used when your service claims to implement IHello
var ReflectTypeIHello = di.GetInterfaceReflectType((*IHello)(nil))

// AddIHelloByObj adds a prebuilt obj
func AddIHelloByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeIHello)
}

// AddSingletonIHello adds a type that implements IHello
func AddSingletonIHello(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeIHello)
}
