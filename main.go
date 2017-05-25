package main

import(
        "fmt"
)

type Categoria struct {
        id              int
        descripcion     string
}

type Riego interface {
        
        
}

type Gotero struct {
        id              int
        lth             float64
        distancia       float64
}

type Aspersor struct {
        id              int
        lth             float64
        distancia       float64
}

type Arbol struct {
        id              int
        descripcion     string
}

type Especie struct {
        id              int
        descripcion     string
}

type Variedad struct {
        id              int
        descripcion     string
}

type Sector struct {
        id              int
        descripcion     string
        arboles         map[Arbol]int
        marco           string
}

type Trabajador struct {
        id              int
        nombre          string
        direccion       string
        rut             string
        telefono        string
        mail            string
        sueldo          float64
        cargo           Cargo
        contrato        Contrato
}

type Cargo struct {
        id              int
        descripcion     string
}

type Contrato struct {
        id              int
        descripcion     string
}

type Bodega struct {
        id              int
        descripcion     string
        detalleProductos        []DetalleProducto        
}

type DetalleProducto struct {
        id              int
        cantidad        float64
        producto        Producto
}

type Producto struct {
        id              int
        descripcion     string
        categoria       CategoriaProducto
}

type CategoriaProducto struct {
        id              int
        descripcion     string
} 

type Administracion struct {
        id              int
        nombre          string
        direccion       string
        rut             string
        telefono        string
        mail            string
        trabajadores    []Trabajador
        bodegas         []Bodega
        sectores        []Sectores
        clientes        []Cliente
        proveedores     []Proveedor        
}

type Cliente struct {
        id              int
        nombre          string
        direccion       string
        rut             string
        telefono        string
        mail            string
}

type Proveedor struct {
        id              int
        nombre          string
        direccion       string
        rut             string
        telefono        string
        mail            string
}

type FacturaCompra struct {
        id              int
        numero          int
        proveedor       Proveedor
        fecha           string
        detalleFacturaCompra  []DetalleFacturaCompra
}

type DetalleFacturaCompra struct {
        id              int
        producto        Producto
        cantidad        float64
        precio          float64  
}       

type FacturaVenta struct {
        id              int
        numero          int
        proveedor       Proveedor
        fecha           string
        detalleFacturaVenta  []DetalleFacturaVenta
}

type DetalleFacturaVenta struct {
        id              int
        producto        Producto
        cantidad        float64
        precio          float64  
}     

func main(){
        fmt.Println("Prueba Parcela")
}
