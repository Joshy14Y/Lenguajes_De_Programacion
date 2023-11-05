import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

class Socio {
    public int numeroSocio;  // Cambiado a propiedad pública
    private String nombre;
    private String direccion;
    private List<Prestamo> prestamos;

    public Socio(int numeroSocio, String nombre, String direccion) {
        this.numeroSocio = numeroSocio;
        this.nombre = nombre;
        this.direccion = direccion;
        this.prestamos = new ArrayList<>();
    }

    public void tomarPrestado(Libro libro, String fecha) {
        if (libro.estaDisponible()) {
            Prestamo prestamo = new Prestamo(libro, this, fecha);
            prestamos.add(prestamo);
            libro.prestar();
        } else {
            System.out.println("El libro no está disponible para préstamo.");
        }
    }

    public int getNumeroLibrosPrestados() {
        return prestamos.size();
    }
}

class Libro {
    private int codigo;
    private String titulo;
    private String autor;
    private boolean disponible;

    public Libro(int codigo, String titulo, String autor) {
        this.codigo = codigo;
        this.titulo = titulo;
        this.autor = autor;
        this.disponible = true;
    }

    public void prestar() {
        disponible = false;
    }

    public void devolver() {
        disponible = true;
    }

    public boolean estaDisponible() {
        return disponible;
    }
}

class Prestamo {
    private Libro libro;
    private Socio socio;
    private String fecha;

    public Prestamo(Libro libro, Socio socio, String fecha) {
        this.libro = libro;
        this.socio = socio;
        this.fecha = fecha;
    }
}

class Biblioteca {
    public static void main(String[] args) {
        Socio socio1 = new Socio(1, "Juan", "Calle A");
        Socio socio2 = new Socio(2, "Maria", "Calle B");

        Libro libro1 = new Libro(101, "El libro 1", "Autor A");
        Libro libro2 = new Libro(102, "El libro 2", "Autor B");
        Libro libro3 = new Libro(103, "El libro 3", "Autor C");

        socio1.tomarPrestado(libro1, "2023-11-05");
        socio1.tomarPrestado(libro2, "2023-11-06");
        socio1.tomarPrestado(libro3, "2023-11-07");

        socio2.tomarPrestado(libro1, "2023-11-08");

        libro1.devolver();

        System.out.println("Libros disponibles:");
        System.out.println("Libro 1: " + libro1.estaDisponible());
        System.out.println("Libro 2: " + libro2.estaDisponible());
        System.out.println("Libro 3: " + libro3.estaDisponible());

        List<Socio> sociosConMasDe3Libros = findSociosConMasDe3Libros(new Socio[]{socio1, socio2});
        System.out.println("Socios con más de 3 libros prestados:");
        for (Socio socio : sociosConMasDe3Libros) {
            System.out.println(socio.numeroSocio + ": " + socio.getNumeroLibrosPrestados());
        }
    }

    public static List<Socio> findSociosConMasDe3Libros(Socio[] socios) {
        return Arrays.stream(socios)
                .filter(socio -> socio.getNumeroLibrosPrestados() > 3)
                .collect(Collectors.toList());
    }
}
