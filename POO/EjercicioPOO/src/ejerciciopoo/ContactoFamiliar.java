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
public class ContactoFamiliar extends Contacto {
    private String parentesco;

    public ContactoFamiliar(Persona persona, String parentesco) {
        super(persona);
        this.parentesco = parentesco;
    }

    /**
     * @return the parentesco
     */
    public String getParentesco() {
        return parentesco;
    }

    /**
     * @param parentesco the parentesco to set
     */
    public void setParentesco(String parentesco) {
        this.parentesco = parentesco;
    }
    
    
    @Override
    public String toString() {
        return "Contacto Familiar: " + getPersona() + ", Parentesco: " + parentesco;
    }

}
