/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 * Hecho por: Noelia Alpízar Torres
 * Objetivo: Cumplir con los ejercicios propuestos para el repositorio.
 * Fecha: 21/05/2024
 */
package ejerciciopoo;

/**
 *Me parece que implementarlo en la creación de eventos y contactos en una buena idea porque
  Facilita la creación de objetos relacionados sin especificar sus clases concretas, mejorando 
  la escalabilidad y flexibilidad. Permite añadir nuevos tipos de contactos y eventos en el futuro 
  sin modificar el código existente.
 */
public abstract class AbstractFactory {
    public abstract Contacto crearContacto(Persona persona);
    public abstract Evento crearEvento(String nombre, String fecha);
 
}
