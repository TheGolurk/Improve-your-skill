package com.company;

public class Alumno{
    //Atributos
    public String Nombre;
    public Float Calificacion;
    public Integer Matricula;

    //Constructor
    public Alumno(String nombre, Float calificacion, Integer matricula) {
        this.Nombre = nombre;
        this.Calificacion = calificacion;
        this.Matricula = matricula;
    }

    //metodo
    public void Saludar(){
        System.out.println("Hola mi nombre es: " + Nombre);
        System.out.println("Saque: " +  Calificacion);
        System.out.println("Mi matricula es: " + Matricula);
    }
}
