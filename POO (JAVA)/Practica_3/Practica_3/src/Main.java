import java.util.ArrayList;

// Singleton Agenda (Eager Singleton)
class Agenda {
    private static final Agenda instance = new Agenda();
    private ArrayList<Object> elementos = new ArrayList<>();

    private Agenda() {
    }

    public static Agenda getInstance() {
        return instance;
    }

    public void agregarElemento(Object elemento) {
        elementos.add(elemento);
    }

    public void eliminarElemento(Object elemento) {
        elementos.remove(elemento);
    }

    @Override
    public String toString() {
        StringBuilder agendaStr = new StringBuilder();
        for (Object elemento : elementos) {
            agendaStr.append(elemento.toString()).append("\n");
        }
        return agendaStr.toString();
    }
}

class ContactoT1 extends Contacto {
    private String direccion;

    public ContactoT1(String nombre, String apellido, String direccion) {
        super(nombre, apellido);
        this.direccion = direccion;
    }

    // Getter y Setter de dirección
    public String getDireccion() {
        return direccion;
    }

    public void setDireccion(String direccion) {
        this.direccion = direccion;
    }

    @Override
    public String toString() {
        return super.toString() + ", Dirección: " + direccion;
    }
}

class ContactoT2 extends Contacto {
    private String telefono;

    public ContactoT2(String nombre, String apellido, String telefono) {
        super(nombre, apellido);
        this.telefono = telefono;
    }

    // Getter y Setter de teléfono
    public String getTelefono() {
        return telefono;
    }

    public void setTelefono(String telefono) {
        this.telefono = telefono;
    }

    @Override
    public String toString() {
        return super.toString() + ", Teléfono: " + telefono;
    }
}

class EventoT1 extends Evento {
    private String tipoEventoT1;

    public EventoT1(String nombreEvento, String tipoEventoT1) {
        super(nombreEvento);
        this.tipoEventoT1 = tipoEventoT1;
    }

    // Getter y Setter del tipo del evento T1
    public String getTipoEventoT1() {
        return tipoEventoT1;
    }

    public void setTipoEventoT1(String tipoEventoT1) {
        this.tipoEventoT1 = tipoEventoT1;
    }

    @Override
    public String toString() {
        return super.toString() + ", Tipo del Evento T1: " + tipoEventoT1;
    }


}

class EventoT2 extends Evento {
    private String tipoEventoT2;

    public EventoT2(String nombreEvento, String tipoEventoT2) {
        super(nombreEvento);
        this.tipoEventoT2 = tipoEventoT2;
    }

    // Getter y Setter del tipo del evento T2
    public String getTipoEventoT2() {
        return tipoEventoT2;
    }

    public void setTipoEventoT2(String tipoEventoT2) {
        this.tipoEventoT2 = tipoEventoT2;
    }

    @Override
    public String toString() {
        return super.toString() + ", Tipo del Evento T2: " + tipoEventoT2;
    }
}

interface ContactoFactory {
    Contacto crearContacto(String nombre, String apellido);
}

interface EventoFactory {
    Evento crearEvento(String nombreEvento);
}

abstract class Contacto {
    private String nombre;
    private String apellido;

    public Contacto(String nombre, String apellido) {
        this.nombre = nombre;
        this.apellido = apellido;
    }

    @Override
    public String toString() {
        return "Nombre: " + nombre + ", Apellido: " + apellido;
    }
}

abstract class Evento {
    private String nombreEvento;

    public Evento(String nombreEvento) {
        this.nombreEvento = nombreEvento;
    }

    @Override
    public String toString() {
        return "Nombre del Evento: " + nombreEvento;
    }
}

class ContactoT1Factory implements ContactoFactory {
    @Override
    public Contacto crearContacto(String nombre, String apellido) {
        return new ContactoT1(nombre, apellido, "Dirección por defecto");
    }
}

class ContactoT2Factory implements ContactoFactory {
    @Override
    public Contacto crearContacto(String nombre, String apellido) {
        return new ContactoT2(nombre, apellido, "Teléfono por defecto");
    }
}

class EventoT1Factory implements EventoFactory {
    @Override
    public Evento crearEvento(String nombreEvento) {
        return new EventoT1(nombreEvento, "Tipo del Evento T1 por defecto");
    }
}

class EventoT2Factory implements EventoFactory {
    @Override
    public Evento crearEvento(String nombreEvento) {
        return new EventoT2(nombreEvento, "Tipo del Evento T2 por defecto");
    }
}

public class Main {
    public static void main(String[] args) {
        // Crear una agenda
        Agenda agenda = Agenda.getInstance();

        // Crear fábricas de Contacto y Evento
        ContactoFactory contactoFactory;
        EventoFactory eventoFactory;

        // Agregar contactos y eventos a la agenda
        contactoFactory = new ContactoT1Factory();
        Contacto contactoT1 = contactoFactory.crearContacto("Juan", "Perez");
        agenda.agregarElemento(contactoT1);

        contactoFactory = new ContactoT2Factory();
        Contacto contactoT2 = contactoFactory.crearContacto("Maria", "Gomez");
        agenda.agregarElemento(contactoT2);

        eventoFactory = new EventoT1Factory();
        Evento eventoT1 = eventoFactory.crearEvento("Fiesta de cumpleaños");
        agenda.agregarElemento(eventoT1);

        eventoFactory = new EventoT2Factory();
        Evento eventoT2 = eventoFactory.crearEvento("Reunión de trabajo");
        agenda.agregarElemento(eventoT2);

        // Mostrar la agenda
        System.out.println("Agenda:");
        System.out.println(agenda);
    }
}
