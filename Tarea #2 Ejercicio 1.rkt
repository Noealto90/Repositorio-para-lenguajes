(define
  listaProductos
  '(("arroz" 8 1800)
    ("frijoles" 5 1200)
    ("azúcar" 6 1100)
    ("cafe" 2 2800)
    ("leche" 9 1200)))

(define (calcularImpuesto factura X)
  (define (calcularImpuestoP producto)
    (if (> (caddr producto) X)
        (* 0.13 (* (cadr producto) (caddr producto)))
        0))
  (apply + (map calcularImpuestoP factura)))



(define (calcularTotal factura)
  (apply + (map (lambda (producto) (* (cadr producto) (caddr producto))) factura)))

(define factura1 '(("arroz" 2 1800)
                   ("frijoles" 3 1200)
                   ("azúcar" 4 1100)))

(define factura2 '(("cafe" 2 2800)
                   ("leche" 3 1200)))

; Casos de prueba
; Si aplica el impuesto el limite es de 1000
(displayln "Factura 1:")
(displayln factura1)
(displayln "Impuesto total a cancelar:")
(displayln (calcularImpuesto factura1 1000))
(displayln "Monto total de la factura sin impuesto:")
(displayln (calcularTotal factura1))
(displayln "Monto total de la factura:")
(displayln (+(calcularImpuesto factura1 1000)(calcularTotal factura1)))
(newline)

; No aplica el impuesto el limite es de 10000
(displayln "Factura 2:")
(displayln factura2)
(displayln "Impuesto total a cancelar:")
(displayln (calcularImpuesto factura2 10000))
(displayln "Monto total de la factura sin impuesto:")
(displayln (calcularTotal factura2))
(displayln "Monto total de la factura:")
(displayln (+(calcularImpuesto factura2 10000)(calcularTotal factura2)))
(newline)

;Si aplica el impuesto el limite es 500
(displayln "Factura 2:")
(displayln factura2)
(displayln "Impuesto total a cancelar:")
(displayln (calcularImpuesto factura2 500))
(displayln "Monto total de la factura sin impuesto:")
(displayln (calcularTotal factura2))
(displayln "Monto total de la factura:")
(displayln (+(calcularImpuesto factura2 500)(calcularTotal factura2)))
(newline)