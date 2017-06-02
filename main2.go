package main

import (
        

)

const (
        VENTA = false
        COMPRA = true
)
//TODO : crear una biblioteca con estas dos funciones
func Round(x, unit float64) float64 {
    return float64(int64(x/unit+0.5)) * unit
}

func RoundInt(x float64) int {
        return int(x+0.5)
}

type Empresa struct {
        id              int
        rut             int
        razonSocial     string
        direccion       string
        comuna          string
        ciudad          string
        giro            string
        telefono        string
        mail            string
}

type Producto struct {
        id              int
        codigo          string
        descripcion     string
        idProveedor     int
}

type Factura struct {
        id              int
        folio           int
        tipo            bool
        fecha           string
        empresa         Empresa
        facturaDetalle  []FacturaDetalle
        descuento       float64
        iva             float64
}

type FacturaDetalle struct {
        producto        Producto
        cantidad        int
        precio          int
        impuesto        float64
        descuento       float64
}

//TODO : Falta incorporar el impuesto específico y el descuento;
func (f *FacturaDetalle) SubTotal() int {
        return RoundInt( (float64)(f.cantidad * f.precio) * (1.0 - f.descuento) * (1.0 + f.impuesto) )
}

func (f *Factura) SubTotal() int {
        subTotal := 0
        for _, fd := range f.facturaDetalle {
                subTotal += fd.SubTotal()
        }
        return subTotal
}

//TODO : descuento expresado como 0.x
func (f *Factura) MontoNeto() int {
        return RoundInt( (float64)(f.SubTotal()) * (1.0 - f.descuento) )
}

//TODO : iva expresado como 0.x
func (f *Factura) Total() int {
        return RoundInt( (float64)(f.MontoNeto()) * (1.0 + f.iva) )
}

func ingresarFacturaCompra() {
        
}

func ingresarProducto() {

}

func ingresarEmpresa() {

}

func main() { 
        
}
