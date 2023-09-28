package models

type KeyValue struct {
	key   string
	value any
}

type KeyValueBuilder struct {
	keyValue *KeyValue
}

func NewKeyValue() *KeyValueBuilder {
	return &KeyValueBuilder{&KeyValue{}}
}

func (kvb *KeyValueBuilder) SetKey(key string) *KeyValueBuilder {
	kvb.keyValue.key = key
	return kvb
}

func (kvb *KeyValueBuilder) SetValue(value any) *KeyValueBuilder {
	kvb.keyValue.value = value
	return kvb
}

func (kvb *KeyValueBuilder) Build() *KeyValue {
	return kvb.keyValue
}

func (kv *KeyValue) GetKey() string {
	return kv.key
}

func (kv *KeyValue) GetValue() any {
	return kv.value
}
