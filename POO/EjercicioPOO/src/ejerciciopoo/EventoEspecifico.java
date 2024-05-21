/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 * Hecho por: Noelia Alpízar Torres
 * Objetivo: Cumplir con los ejercicios propuestos para el repositorio.
 * Fecha: 21/05/2024
 */
package ejerciciopoo;

/**
 *
 * @author HP
 */
public class EventoEspecifico extends Evento {
    private String lugar;
    private String descripcion;
    
    public EventoEspecifico(String nombre, String fecha, String lugar, String descripcion) {
        super(nombre, fecha);
        this.lugar = lugar;
        this.descripcion = descripcion;
    }

    /**
     * @return the lugar
     */
    public String getLugar() {
        return lugar;
    }

    /**
     * @param lugar the lugar to set
     */
    public void setLugar(String lugar) {
        this.lugar = lugar;
    }

    /**
     * @return the descripcion
     */
    public String getDescripcion() {
        return descripcion;
    }

    /**
     * @param descripcion the descripcion to set
     */
    public void setDescripcion(String descripcion) {
        this.descripcion = descripcion;
    }
    
    @Override
    public String toString() {
        return "Evento Especifico: " + getNombre() + ", Fecha: " + getFecha() + ", Lugar: " + lugar + ", Descripción: " + descripcion;
    }
    
}
