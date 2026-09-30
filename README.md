# CulturaGRX

## Problema

Este problema lo tienen muchos ciudadanos de la ciudad de Granada, al igual que muchos turistas que vienen a visitar la ciudad y provincia. Cuando quieren, en su tiempo libre, vacaciones o fines de semana, o en el caso de los turistas en su visita, no pueden ni saben descubrir ni visitar muchísimos espacios culturales y patrimoniales, ya que la información de la que disponemos de su existencia es prácticamente nula. Además, tampoco saben sus horarios de visita, la historia que tienen detrás esos espacios, etc.

También tienen la necesidad de organizar las visitas cuando se tiene poco tiempo y unos intereses concretos, según los horarios.

Es un problema porque muchos de estos edificios se encuentran sin visitas ya que muchas personas no asisten porque no saben ni de su existencia. Además, el trabajo y esfuerzo de muchas personas se ve gravemente afectado e incluso en peligro.

## Fuentes de datos

El problema necesita información sobre los diferentes bienes culturales y patrimoniales de Granada y provincia, que son joyas que están sin descubrir y por falta de información no conocemos ni están al alcance de nosotros. De diferentes fuentes abiertas podremos obtener información tanto de sus ubicaciones, historia, etc.

La primera que voy a usar es esta fuente de datos para el proyecto, el primero es de la web de la Junta de Andalucía y su Instituto Andaluz de Patrimonio Histórico, esta fuente la voy a utilizar para la parte de qué espacios existen y dónde están,  en el cual dentro de esta web: [portal de datos abiertos de la Junta](https://www.juntadeandalucia.es/datosabiertos/portal/dataset/patrimonio-inmueble-de-andalucia) nos encontramos con una API con formato XML, CSV y JSON para aportar id, codigo, denominacion, provincia, municipio y caracterizacion. En un ejemplo de la API (json): [API en formato JSON](https://www.juntadeandalucia.es/datosabiertos/portal/iaph/dataset/bien/inmueble?page=0&rows=31000&format=json)

```json
 {
        "id":"347310",
        "codigo":"01180020009",
        "denominacion":"El Hacho",
        "caracterizacion":"Arqueológica"},
      {
        "caracterizacion":"Arquitectónica",
        "codigo":"01180030001",
        "denominacion":"Iglesia de la Encarnación",
        "id":"9232",
        "municipio":"Albolote",
        "provincia":"Granada"},
      {
        "caracterizacion":"Arquitectónica, Etnológica",
        "codigo":"01180030003",
        "denominacion":"Toro Osborne XI",
        "id":"499",
        "municipio":"Albolote",
        "provincia":"Granada"},

```

Aquí podemos extraer los datos que he comentado antes, su identificación, su código (que como podemos observar todos los de Granada empiezan por 0118), denominación, el municipio y la caracterización.

Además de por cada bien patrimonial, tenemos dentro de su página personalizada por ejemplo de la iglesia de la Encarnación de Albolote, con la siguiente URL:  [ficha de la Iglesia de la Encarnación](https://guiadigital.iaph.es/bien/inmueble/9232/granada/albolote/iglesia-de-la-encarnacion) dentro de la misma podemos descargar arriba a la derecha los datos abierto de cada uno con (.jsonld), por ejemplo:

[ficha-inmueble-9232.jsonld](docs/ficha-inmueble-9232.jsonld) 

Aquí dentro podemos extraer las coordenadas, la protección, el periodo y la descripción.

Para poder acceder a estos datos abiertos, sería necesario entrar en la URL específica de cada bien patrimonial, que se ubica un botón justo en la esquina derecha llamado Descargar datos abiertos, esto lo he conseguido buscando en las herramientas del desarrollador, dentro de Network y escribiendo dentro de Filter el id del bien que quieras encontrar y pulsando en Headers encuentras la URL pertinente, por ejemplo: `https://guiadigital.iaph.es/api/1.0/bien/inmueble/enriquecido/9232` , como vemos el último número (id), es el que hacer referencia a uno concreto, pero al abrirla directamente da error,  porque la API exige un token de acceso. La web de la ficha llama a esa URL enviando un token; por eso en las herramientas de desarrollador la petición funciona, pero si abres la URL tú solo, sin token, da error. Apareciendo el siguiente mensaje: 

```
<ams:fault xmlns:ams="http://wso2.org/apimanager/security">
<ams:code>900902</ams:code>
<ams:message>Missing Credentials</ams:message>
<ams:description>Required OAuth credentials not provided. Make sure your API invocation call has a header: "Authorization: Bearer ACCESS_TOKEN"</ams:description>
</ams:fault>
```

En limitaciones encontramos en primer lugar con que en el IAPH no hay información sobre el estado de conservación de cada bien, tampoco contempla horarios y tiene dos fallos de calidad: Latitud y longitud están intercambiadas. En la iglesia de la Encarnación de Albolote la marca en "latitud_s": -3.657 y "longitud_s": 37.23, y en cambio Albolote está en latitud 37.23 y longitud -3.66. Como vemos estos valores están puestos al revés. Cuando vayamos a usar las coordenadas tenemos que tener esto en cuenta. El segundo fallo de calidad que he encontrado es que la bibliografía a veces no corresponde ya que en este ejemplo aparece un libro sobre un yacimiento de la Edad del Bronce en Purullena, que no tiene nada que ver con esta iglesia. 
También descubrí otra limitación dentro de esta página, esta es que muchos registros dentro de Granada y provincia no cuentan con su municipio ni provincia, por lo que va a haber que identificarlos por el principio del código de los bienes de Granada (0118). 

La segunda fuente de datos que voy a usar es una dentro de los datos abiertos de la Junta de Andalucía, llamada Museos y Colecciones Museográficas de Andalucía: [Museos y Colecciones Museográficas de Andalucía](https://www.juntadeandalucia.es/datosabiertos/portal/dataset/museos-de-andalucia) dentro de ella para poder extraer nuestros datos necesarios referentes a los horarios de apertura, lugares en los que se encuentran vamos a introducirnos en el formato JSON: [JSON completo de museos](https://datos.juntadeandalucia.es/api/v0/museums/all?format=json) y al abrirla se descarga el archivo completo. 

Un ejemplo de extracción de datos con un museo sería el siguiente: 

[museo-ejemplo.json](docs/museo-ejemplo.json) 

```json
{
    "id": 2702,
    "name": "Parque de las Ciencias",
    "location": "Granada",
    "postcode": 18006,
    "latitude": "37,15949",
    "municipality": "Granada",
    "observations": "El Parque de las Ciencias tiene por objetivo promover el inter\u00e9s por la cultura cient\u00edfica. Dispone de salas expositivas, donde se facilita la comprensi\u00f3n de fen\u00f3menos cient\u00edficos, tecnol\u00f3gicos, medioambientales y en particular sobre f\u00edsica, astronom\u00eda, percepci\u00f3n, etc., a trav\u00e9s de la manipulaci\u00f3n de diferentes m\u00f3dulos bajo el principio de la interactividad. El centro cuenta con \u00e1reas exteriores ajardinadas provistas de contenido expositivo. En ellas, se puede disfrutar de recorridos bot\u00e1nicos, m\u00e1quinas e ingenios, laberinto vegetal, mariposarium, almazara, paseo de los pavimentos, etc. El edificio est\u00e1 integrado por dos grandes alas unidas por un hall de cristal y est\u00e1 dotado con servicio de cafeter\u00eda, tienda, planetario, observatorio astron\u00f3mico, planetario infantil, sala Explora para ni\u00f1os de 3 a 7 a\u00f1os y sala dedicada a exposiciones temporales.",
    "address": "Avenida Avenida de la Ciencia s/n, Consorcio Parque de las Ciencias, 18006 Granada",
    "opening_hours": "Martes a s\u00e1bado de 10:00 a 19:00. domingos y festivos de 10:00 a 15:00. de Martes a S\u00e1bado de 10:00 a 19:00. Domingo y festivos de 10:00 a 15:00.",
    "web": "http://www.parqueciencias.com",
    "province": "Granada",
    "longitude": "-3,609117",
    "unit_type": "Museo",
    "state": "ABI",
    "phone": "958 13 19 00",
    "fax": "958 13 35 82",
    "email": "info@parqueciencias.com"
}

```

Los campos que voy a usar aquí son sobre todo los horarios, días festivos abiertos u obviamente la localización de cada uno.

Como he podido observar, en el apartado del esfuerzo necesario para extraer los datos, aquí si va a haber que interpretarlos, ya que vienen como un texto libre y no como datos estructurados, por ejemplo en el opening hours aparece de martes a sábado en vez de cada día de la semana con su horario correspondiente, también el texto no es uniforme ya que repite el horario escrito de diferentes formas.
Además también las coordenadas tanto en longitud como en latitud aparecen con coma decimal, así que también habrá que reconvertirlas a número. 

Una de las limitaciones que encuentro es que los horarios y días de apertura solamente están para los 88 museos disponibles en la provincia de Granada, cosas que para los bienes culturales no, también es verdad que algunos de ellos no tienen horarios porque no están abiertos al público, solamente se pueden visitar desde fuera.


## Juego de rol

![Ficha del juego de rol](docs/ficha.jpg)

## Documentación

- [Configuración del entorno](docs/configuracion.md)
