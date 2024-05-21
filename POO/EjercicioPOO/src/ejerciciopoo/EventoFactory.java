/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 * Hecho por: Noelia Alpízar Torres
 * Objetivo: Cumplir con los ejercicios propuestos para el repositorio.
 * Fecha: 21/05/2024
 */
package ejerciciopoo;

/**
 * Fábrica concreta para la creación de Eventos.
 */
public class EventoFactory extends AbstractFactory {
    @Override
    public Contacto crearContacto(Persona persona) {
        throw new UnsupportedOperationException("Esta fábrica solo crea eventos.");
    }

    @Override
    public Evento crearEvento(String nombre, String fecha) {
        return new Evento(nombre, fecha);
    }
    
}
