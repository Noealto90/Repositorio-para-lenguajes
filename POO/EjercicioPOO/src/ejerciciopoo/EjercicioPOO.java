/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 * Hecho por: Noelia Alpízar Torres
 * Objetivo: Cumplir con los ejercicios propuestos para el repositorio.
 * Fecha: 21/05/2024
 */
package ejerciciopoo;

import java.util.List;

/**
 *
 * @author HP
 */
public class EjercicioPOO {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
       // Crear algunas personas
        Persona p1 = new Persona("Juan", "Perez", "juan.perez@gmail.com");
        Persona p2 = new Persona("Maria", "Gomez", "maria.gomez@gmail.com");
        Persona p3 = new Persona("Carlos", "Lopez", "carlos.lopez@gmail.com");

        // Crear fábricas
        AbstractFactory contactoFactory = new ContactoFactory();
        AbstractFactory eventoFactory = new EventoFactory();

        // Crear contactos usando la fábrica
        Contacto c1 = contactoFactory.crearContacto(p1);
        ContactoFamiliar cf1 = new ContactoFamiliar(p2, "Hermana");
        ContactoEmpresarial ce1 = new ContactoEmpresarial(p3, "TechCorp", "Gerente");

        // Crear eventos usando la fábrica
        Evento e1 = eventoFactory.crearEvento("Reunion", "2023-06-15");
        EventoEspecifico ee1 = new EventoEspecifico("Conferencia", "2023-07-20", "Auditorio", "Conferencia anual de tecnologia");

        // Obtener la instancia única de la agenda y agregar elementos
        Agenda agenda = Agenda.getInstancia();
        agenda.agregarElemento(c1);
        agenda.agregarElemento(cf1);
        agenda.agregarElemento(ce1);
        agenda.agregarElemento(e1);
        agenda.agregarElemento(ee1);

        // Imprimir la agenda completa
        System.out.println("Agenda completa:");
        System.out.println(agenda);

        // Filtrar contactos usando programación funcional
        List<Contacto> contactos = agenda.filtrarContactos();
        System.out.println("\nContactos:");
        contactos.forEach(System.out::println);

        // Filtrar eventos usando programación funcional
        List<Evento> eventos = agenda.filtrarEventos();
        System.out.println("\nEventos:");
        eventos.forEach(System.out::println);
    }
    
}
