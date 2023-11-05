import java.util.ArrayList;
import java.util.List;

class Documento {
    private List<Hoja> hojas = new ArrayList<>();

    public void agregarHoja(Hoja hoja) {
        hojas.add(hoja);
    }

    public void eliminarHoja(Hoja hoja) {
        hojas.remove(hoja);
    }
}

class Hoja {
    private List<ObjetoRepresentable> objetos = new ArrayList<>();

    public void agregarObjeto(ObjetoRepresentable objeto) {
        objetos.add(objeto);
    }

    public void eliminarObjeto(ObjetoRepresentable objeto) {
        objetos.remove(objeto);
    }
}

abstract class ObjetoRepresentable {
    private int x;
    private int y;

    public int getX() {
        return x;
    }

    public void setX(int x) {
        this.x = x;
    }

    public int getY() {
        return y;
    }

    public void setY(int y) {
        this.y = y;
    }
}

class Grupo extends ObjetoRepresentable {
    private List<ObjetoRepresentable> objetosEnGrupo = new ArrayList<>();

    public void agregarObjetoEnGrupo(ObjetoRepresentable objeto) {
        objetosEnGrupo.add(objeto);
    }

    public void eliminarObjetoDeGrupo(ObjetoRepresentable objeto) {
        objetosEnGrupo.remove(objeto);
    }
}

class Circulo extends ObjetoRepresentable {
    private int radio;

    public Circulo(int x, int y, int radio) {
        this.setX(x);
        this.setY(y);
        this.radio = radio;
    }

    // Métodos específicos para círculos
}

class Rectangulo extends ObjetoRepresentable {
    private int ancho;
    private int alto;

    public Rectangulo(int x, int y, int ancho, int alto) {
        this.setX(x);
        this.setY(y);
        this.ancho = ancho;
        this.alto = alto;
    }

    // Métodos específicos para rectángulos
}

public class Main {
    public static void main(String[] args) {
        Documento documento = new Documento();
        Hoja hoja = new Hoja();
        Grupo grupo = new Grupo();

        Circulo circulo = new Circulo(50, 50, 30);
        Rectangulo rectangulo = new Rectangulo(100, 100, 60, 40);

        grupo.agregarObjetoEnGrupo(circulo);
        grupo.agregarObjetoEnGrupo(rectangulo);

        hoja.agregarObjeto(grupo);
        documento.agregarHoja(hoja);
    }
}