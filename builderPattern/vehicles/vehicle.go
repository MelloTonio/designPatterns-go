package vehicles

// O padrão builder separa a criação de um objeto em várias partes diferentes
// desse modo tornando sua inicialização mais simples e intuitiva
// neste exemplo separamos a criação de um carro em varias etapas
// util para criar objetos complexos
type IBuilder interface {
	SetBrand(brand string) *CarBuilder
	SetLicensePlate(licensePlate string) *CarBuilder
	SetColor(color string) *CarBuilder
	SetSpeed(speed int) *CarBuilder
	Get() *Car
}

type CarBuilder struct {
	Brand        string
	LicensePlate string
	Color        string
	Speed        int
	Type         string
}

func (cb *CarBuilder) SetBrand(brand string) *CarBuilder {
	cb.Brand = brand

	return cb
}

func (cb *CarBuilder) SetLicensePlate(licensePlate string) *CarBuilder {
	cb.LicensePlate = licensePlate

	return cb
}

func (cb *CarBuilder) SetColor(color string) *CarBuilder {
	cb.Color = color

	return cb
}

func (cb *CarBuilder) SetSpeed(speed int) *CarBuilder {
	cb.Speed = speed

	return cb
}

func (cb *CarBuilder) Get() *Car {
	return &Car{
		Brand:        cb.Brand,
		Color:        cb.Color,
		LicensePlate: cb.LicensePlate,
		Speed:        cb.Speed,
		Type:         "CAR",
	}
}

func New() CarBuilder {
	return CarBuilder{}
}
