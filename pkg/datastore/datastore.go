package datastore

// Datastore -
type Datastore interface {
	List(castTo interface{}) error
	Get(key string, castTo interface{}) error
	Store(item interface{}) error
}
