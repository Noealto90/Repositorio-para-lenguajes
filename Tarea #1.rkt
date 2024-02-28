;;Tarea #1 de lenguajes de programación
;;Fecha: 25/02/2024
;;Hecho por: Noelia Alpízar Torres
;;Objetivo: Realizar los ejercicios propuestos por el profesor

;;Ejercicio #1
;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(define (int Cap I N)
  (define (int capital I N)
    (cond ((= N 0) capital)
          ((= N 1) (+ capital (* capital I)))
          (else (int (+ capital (* capital I)) I (- N 1)))))
  
  (int Cap I N))




;;Ejercicio #6
;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(define (merge lista1 lista2)
  (cond
    ((null? lista1) lista2)
    ((null? lista2) lista1)
    ((< (car lista1) (car lista2))
     (cons (car lista1) (merge (cdr lista1) lista2)))
    (else
     (cons (car lista2) (merge lista1 (cdr lista2))))))


;;Ejercicio #8
;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(define (sub-conjunto? lista1 lista2)
  (cond
    ((null? lista1) #t) ; Si lista1 está vacía, entonces todos sus elementos están en lista2
    ((member (car lista1) lista2) (sub-conjunto? (cdr lista1) lista2))
    (else #f)))

;;Ejercicio #9
;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(define (eliminar_elemento E L)
  (define (eliminar x)
    (if (equal? x E)
        '()
        (list x)))
  (apply append (map eliminar L)))