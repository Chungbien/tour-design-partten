package FactoryMethod

type ReligonFactory interface {
  Create(religon string, logo string, size int, o string)
}