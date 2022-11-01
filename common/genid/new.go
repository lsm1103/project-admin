package genid

import "project-admin/common/genid/generator"

// NewGeneratorData
func NewGeneratorData() (ret *generator.GeneratorData) {
	var (
		data = new(generator.GeneratorData)
	)
	data.GeneratorBankID()
	data.GeneratorAddress()
	data.GeneratorEmail()
	data.GeneratorIDCart()
	data.GeneratorName()
	data.GeneratorPhone()
	ret = data
	return
}
