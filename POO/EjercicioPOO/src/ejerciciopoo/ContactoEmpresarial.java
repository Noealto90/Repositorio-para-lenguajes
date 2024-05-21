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
public class ContactoEmpresarial extends Contacto{
    private String empresa;
    private String cargo;
    
    
    public ContactoEmpresarial(Persona persona, String empresa, String cargo) {
        super(persona);
        this.empresa = empresa;
        this.cargo = cargo;
    }

    /**
     * @return the empresa
     */
    public String getEmpresa() {
        return empresa;
    }

    /**
     * @param empresa the empresa to set
     */
    public void setEmpresa(String empresa) {
        this.empresa = empresa;
    }

    /**
     * @return the cargo
     */
    public String getCargo() {
        return cargo;
    }

    /**
     * @param cargo the cargo to set
     */
    public void setCargo(String cargo) {
        this.cargo = cargo;
    }
    
    
    @Override
    public String toString() {
        return "Contacto Empresarial: " + getPersona() + ", Empresa: " + empresa + ", Cargo: " + cargo;
    }
}
