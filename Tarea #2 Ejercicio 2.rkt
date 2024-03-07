(define (sub_cadenas subcadena lista)
  (define (contiene-subcadena? cadena)
    (cond
      ((string? cadena)
       (let loop ((i 0))
         (cond
           ((>= i (- (string-length cadena) (string-length subcadena))) #f)
           ((equal? (substring cadena i (+ i (string-length subcadena))) subcadena) #t)
           (else (loop (+ i 1))))))
      (else #f)))

  (filter contiene-subcadena? lista))


;(sub_cadenas "la" '("la casa" "el perro" "pintando la cerca"))