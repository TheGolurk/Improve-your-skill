package com.company;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        //lista dinamica
        List<Alumno> alumnos = new ArrayList<>();
        Scanner leer = new Scanner(System.in);

        String nom = ""; float cal = 0f; int mat = 0;
        for (int i = 0; i <= 1; i++){
            nom = leer.next();
            cal = leer.nextFloat();
            mat = leer.nextInt();

            Alumno al1 = new Alumno(nom, cal, mat);
            alumnos.add(al1); // <- lista
            al1.Saludar();
        }

        for (int indice = 0; indice <= alumnos.size() - 1; indice++){
            System.out.println( alumnos.get(indice).Nombre  );
            alumnos.get(indice).Saludar();
        }

	    //Alumnos
        //Alumno alumno1 = new Alumno("Andrea", 9.9f, 111);
        //alumno1.Saludar();
    }
}
