package Secure

import (
	"fmt"
	"strings"

	"github.com/6uf/Encrypt"
	"github.com/jaypipes/ghw"
)

const (
	DecodeToken = "c0ef529e570411ed9b6a0242ac120002"
)

func EncryptData() (string, error) {
	opt := ghw.WithDisableWarnings()
	bios, err := ghw.BIOS(opt)
	if err != nil {
		return "", err
	}
	gpu, err := ghw.GPU(opt)
	if err != nil {
		return "", err
	}
	var Data string
	if len(gpu.GraphicsCards) > 0 {
		Data = "-" + gpu.GraphicsCards[0].DeviceInfo.Product.ID
	}
	baseboard, err := ghw.Baseboard(opt)
	if err != nil {
		return "", err
	}
	product, err := ghw.Product(opt)
	if err != nil {
		return "", err
	}
	return Encrypt.Encode(fmt.Sprintf("\n\n %v-%v \n\n %v-%v-%v-%v \n\n %v-%v%v \n |k__wdmmWOowkdflPP{wdkWSWS23FD|",
		bios.Vendor, baseboard.Vendor,
		baseboard.SerialNumber, product.SerialNumber, product.UUID, product.Name,
		baseboard.Product, Encrypt.Encode(baseboard.SerialNumber, strings.Replace(product.UUID, "-", "", -1)), Data), "c0ef529e570411ed9b6a0242ac120002"), nil
}
