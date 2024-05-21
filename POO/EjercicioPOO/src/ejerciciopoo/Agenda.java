/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 * Hecho por: Noelia Alpízar Torres
 * Objetivo: Cumplir con los ejercicios propuestos para el repositorio.
 * Fecha: 21/05/2024
 */
package ejerciciopoo;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

/**
 * JUSTIFICACIÓN
 * Personalmente creo que la clase Agenda es un buen lugar para aplicar el patrón Singleton porque representa 
  una entidad única que debería ser accesible de manera global en toda la aplicación.
  Usar un Singleton garantiza que no se creen múltiples instancias de la agenda, evitando inconsistencias.
 */
public class Agenda {
    
    // Instancia única de la clase Agenda
    private static Agenda instancia;

    // Lista de elementos que pueden ser contactos o eventos
    private final List<Object> elementos;

    // Constructor privado para evitar la creación de instancias
    private Agenda() {
        this.elementos = new ArrayList<>();
    }

    // Método estático para obtener la instancia única de la clase Agenda
    public static Agenda getInstancia() {
        if (instancia == null) {
            instancia = new Agenda();
        }
        return instancia;
    }

    public void agregarElemento(Object elemento) {
        elementos.add(elemento);
    }

    public void eliminarElemento(Object elemento) {
        elementos.remove(elemento);
    }

    public void modificarElemento(int index, Object nuevoElemento) {
        if (index >= 0 && index < elementos.size()) {
            elementos.set(index, nuevoElemento);
        }
    }

    public List<Contacto> filtrarContactos() {
        return elementos.stream()
                .filter(elemento -> elemento instanceof Contacto)
                .map(elemento -> (Contacto) elemento)
                .collect(Collectors.toList());
    }

    public List<Evento> filtrarEventos() {
        return elementos.stream()
                .filter(elemento -> elemento instanceof Evento)
                .map(elemento -> (Evento) elemento)
                .collect(Collectors.toList());
    }

    @Override
    public String toString() {
        return elementos.stream()
                .map(Object::toString)
                .collect(Collectors.joining("\n"));
    }
    
}

/*
*DIFERENCIAS ENTRE EAGER SINGLETON Y LAZY SINGLETON

-Eager Singleton:
En este tipo, la instancia de la clase singleton se crea inmediatamente durante el proceso de carga 
de la clase, independientemente de si es necesaria o no. Esto garantiza que la instancia esté disponible 
de inmediato.

-Lazy Singleton: 
En este tipo, la instancia de la clase singleton se crea solo cuando se requiere por primera vez. 
Esto significa que la instancia no se crea durante el proceso de carga de la clase, lo cual puede ser 
beneficioso si la creación de la instancia es costosa en términos de recursos y no siempre es necesaria.

-Diferencias en la creación de la instancia: 
Lazy: La instancia se crea solo cuando se solicita por primera vez.
Eager: La instancia se crea tan pronto como la clase se carga en la memoria.

-Diferencias en el uso de recursos:
Lazy: Puede ahorrar recursos si la instancia no es necesaria durante la ejecución del programa.
Eager: Puede consumir recursos innecesariamente si la instancia no se utiliza.

--Diferencias en la implementación:
Lazy: Requiere una lógica adicional para verificar si la instancia ya ha sido creada.
Eager: Es más simple de implementar ya que la instancia se crea inmediatamente.

*ELECCIÓN
Se ha usado Lazy Singleton porque es más eficiente en términos de recursos si la instancia 
no es necesaria de inmediato. En este caso, la instancia de la agenda solo se crea cuando 
se llama a getInstancia() por primera vez.
*/